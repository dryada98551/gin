package configs

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	GLogging "main/internal/lib/gcplog"
	ginset "main/internal/lib/ginset"
	libhttp "main/internal/lib/http"

	pb "main/internal/lib/jwt"

	"github.com/hellofresh/health-go/v5"
	healthHttp "github.com/hellofresh/health-go/v5/checks/http"
	healthPg "github.com/hellofresh/health-go/v5/checks/postgres"
	healthRedis "github.com/hellofresh/health-go/v5/checks/redis"
)

type InitParmsConfig struct {
	EnvVar         EnvVar
	ReadStaticFile ReadStaticFile
	GinSet         GinSet
	Redis          Redis
	GrpcConn       GrpcConn
	HttpUrl        libhttp.HttpUrl
	PubSub         PubSub
}

var Global InitParmsConfig

func (param InitParmsConfig) JsonSetting() InitParmsConfig {
	GIN_MODE := os.Getenv("GIN_MODE")

	switch GIN_MODE {
	case "release":
		gin.SetMode(gin.ReleaseMode)
		param.HttpUrl = libhttp.HttpUrlRelease
	case "uat":
		gin.SetMode(gin.TestMode)
		param.HttpUrl = libhttp.HttpUrlUat
	case "debug":
		gin.SetMode(gin.DebugMode)
		param.HttpUrl = libhttp.HttpUrlDebug
	default:
		// 如果没有设置环境变量，默认为 debug 模式
		gin.SetMode(gin.DebugMode)
		GIN_MODE = "debug"
		param.HttpUrl = libhttp.HttpUrlDebug
	}

	path, _ := os.Getwd()
	file, err := os.Open(path + "/config/config-" + GIN_MODE + ".json")
	if err != nil {
		GLogging.ERROR("config file not fund")
		os.Exit(1) // 解碼出錯，提前終止函數
	}
	defer file.Close()

	// 創建 JSON 解碼器
	decoder := json.NewDecoder(file)

	// 解碼 JSON
	err = decoder.Decode(&param.EnvVar)
	if err != nil {
		GLogging.ERROR("Error decoding JSON")
		os.Exit(1) // 解碼出錯，提前終止函數
	}
	return param
}

func (param InitParmsConfig) GinSetting() InitParmsConfig {
	param.GinSet.Engine = gin.Default()
	param.GinSet.Engine.Use(ginset.CORSMiddleware())
	// 使用domainCheckMiddleware中間件
	if os.Getenv("GIN_MODE") == "release" {
		param.GinSet.Engine.Use(ginset.DomainCheckMiddleware(param.EnvVar.GIN.HOST))
	}

	return param
}

func (param InitParmsConfig) PubSubSetting() InitParmsConfig {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, param.EnvVar.PUBSUB.PROJECTID)
	if err != nil {
		GLogging.ERROR("Error connect pubsub")
		os.Exit(1) // 解碼出錯，提前終止函數
	}
	topic := client.Topic(param.EnvVar.PUBSUB.TOPICID)
	// 檢查主題是否存在
	exists, err := topic.Exists(ctx)
	if err != nil {
		GLogging.ERROR("Error checking if topic exists: " + err.Error())
		os.Exit(1) // 解碼出錯，提前終止函數
	}
	if !exists {
		GLogging.ERROR("Topic " + param.EnvVar.PUBSUB.TOPICID + " does not exist")
		os.Exit(1) // 解碼出錯，提前終止函數
	}

	param.PubSub.Topic = topic

	subscription := client.Subscription(param.EnvVar.PUBSUB.SUBSCRIPTIONID)
	exists, err = subscription.Exists(ctx)
	if err != nil {
		GLogging.ERROR("Failed to check if subscription exists: " + err.Error())
		os.Exit(1) // 解碼出錯，提前終止函數
	}
	if !exists {
		GLogging.ERROR("Subscription " + param.EnvVar.PUBSUB.SUBSCRIPTIONID + " does not exist")
		os.Exit(1) // 解碼出錯，提前終止函數
	} else {
		GLogging.ERROR("Subscription " + param.EnvVar.PUBSUB.SUBSCRIPTIONID + " already exists")
	}

	param.PubSub.Subscription = subscription

	return param
}

func (param InitParmsConfig) PostgreSQL() InitParmsConfig {
	passwd, err := base64.StdEncoding.DecodeString(param.EnvVar.DB.PASSWORD)
	if err != nil {
		GLogging.ERROR("Error decoding string: " + err.Error())
		os.Exit(1) // 解碼出錯，提前終止函數
	}

	// 初始化數據庫連接參數
	dsn := "user=" + param.EnvVar.DB.USER
	dsn = dsn + " password=" + string(passwd)
	dsn = dsn + " dbname=" + param.EnvVar.DB.DBNAME
	dsn = dsn + " host=" + param.EnvVar.DB.HOST
	dsn = dsn + " port=" + param.EnvVar.DB.PORT + " sslmode=disable"

	// 連接到 PostgreSQL 數據庫
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		GLogging.ERROR("failed to connect database")
		os.Exit(1) // 解碼出錯，提前終止函數
	}
	// sqlDB, err := db.DB()
	// defer sqlDB.Close() // 確保在函數結束後關閉數據庫連接

	param.ReadStaticFile.DB = db
	return param
}

func (param InitParmsConfig) RedisSet() InitParmsConfig {
	// 靜態文件路由
	param.Redis.Client = redis.NewClient(&redis.Options{
		Addr:     param.EnvVar.REDIS.HOST + ":" + param.EnvVar.REDIS.PORT, // Redis 地址
		Password: param.EnvVar.REDIS.PASSWORD,                             // 密碼，如果沒有設置則為空
		DB:       param.EnvVar.REDIS.DB,                                   // 使用的 DB 編號
	})

	var ctx = context.Background()

	// 執行一個簡單的命令來測試連接
	_, err := param.Redis.Client.Ping(ctx).Result()
	if err != nil {
		GLogging.ERROR("failed to connect redis")
		os.Exit(1) // 解碼出錯，提前終止函數
	}
	return param
}

func (param InitParmsConfig) GrpcJwtSet() InitParmsConfig {

	conn, err := grpc.Dial(param.EnvVar.JWT.GRPCHOST+":"+param.EnvVar.JWT.GRPCPORT, grpc.WithInsecure())
	if err != nil {
		GLogging.ERROR("failed to connect jwt grpc")
	}
	//defer conn.Close()
	param.GrpcConn.Client = pb.NewJwtServiceClient(conn)

	return param
}


func (param InitParmsConfig) LivenessSet() {
	param.GinSet.Engine.GET("/_liveness", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})
}

func (param InitParmsConfig) HealthCheckSet() {
	h, _ := health.New(health.WithSystemInfo())

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		param.EnvVar.DB.USER,
		param.EnvVar.DB.PASSWORD,
		param.EnvVar.DB.HOST,
		param.EnvVar.DB.PORT,
		param.EnvVar.DB.DBNAME)

	h.Register(health.Config{
		Name:      "postgres-check",
		Timeout:   time.Second * 5,
		SkipOnErr: true,
		Check: healthPg.New(healthPg.Config{
			DSN: dsn,
		}),
	})

	dsn = fmt.Sprintf("%s:%s",
		param.EnvVar.REDIS.HOST,
		param.EnvVar.REDIS.PORT)

	h.Register(health.Config{
		Name:      "redis-check",
		Timeout:   time.Second * 5,
		SkipOnErr: true,
		Check: healthRedis.New(healthRedis.Config{
			DSN: dsn,
		}),
	})

	for index, url := range param.EnvVar.HealthUrl {
		h.Register(health.Config{
			Name:      "http-check-" + string(index),
			Timeout:   time.Second * 5,
			SkipOnErr: true,
			Check: healthHttp.New(healthHttp.Config{
				URL: url,
			}),
		})
	}

	param.GinSet.Engine.GET("/healthz", gin.WrapH(h.Handler()))
}

// Init initializes the system with the specified configurations.
// api: if true, the API component will be initialized.
// task: if true, the task scheduler will be initialized.
// corn: if true, the corn job processor will be initialized.
func Init(api bool, task bool, corn bool) {
	if api {
		Global = Global.JsonSetting()
		Global = Global.GinSetting()
		Global = Global.PostgreSQL()
		Global = Global.RedisSet()
		Global = Global.GrpcJwtSet()
		Global.LivenessSet()
		Global.HealthCheckSet()
	}
	if task {
		Global = Global.JsonSetting()
		Global = Global.PubSubSetting()
		Global = Global.PostgreSQL()
		Global = Global.RedisSet()
	}
	if corn {
		Global = Global.JsonSetting()
		Global = Global.PostgreSQL()
		Global = Global.RedisSet()
	}
}
