package configs

import (
	"cloud.google.com/go/pubsub"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	pb "main/internal/lib/jwt"
)

type GinSet struct {
	Engine *gin.Engine
}

type PubSub struct {
	Topic        *pubsub.Topic
	Subscription *pubsub.Subscription
}

type ReadStaticFile struct {
	DB *gorm.DB
}

type Redis struct {
	Client *redis.Client
}

type GrpcConn struct {
	Client pb.JwtServiceClient
}

type EnvVar struct {
	PROJECT     EnvPROJECT  `json:"PROJECT"`
	JWT         EnvJWT      `json:"JWT"`
	GIN         EnvGIN      `json:"GIN"`
	PUBSUB      EnvPUBSUB   `json:"PUBSUB"`
	DB          EnvDB       `json:"DB"`
	REDIS       EnvREDIS    `json:"REDIS"`
	GAME        []EnvGAME   `json:"GAME"`
	HealthUrl   []string    `json:"HealthUrl"`
	APIServices APIServices `json:"APIServices"`
	ONEALL      EnvONEALL   `json:"ONEALL"`
	Twilio      EnvTwilio   `json:"Twilio"`
	WEBOWNER    EnvWEBOWNER `json:"WEBOWNER"`
	CRON        EnvCRON     `json:"CRON"`
}

type EnvPROJECT struct {
	NAME string `json:"NAME"`
}

type EnvJWT struct {
	GRPCHOST    string `json:"GRPCHOST"`
	GRPCPORT    string `json:"GRPCPORT"`
	JWTREDISKEY string `json:"JWTREDISKEY"`
}

type EnvGIN struct {
	HOST     string `json:"HOST"`
	PORT     string `json:"PORT"`
	DOMAIN   string `json:"DOMAIN"`
	BASEPATH string `json:"BASEPATH"`
}

type EnvPUBSUB struct {
	PROJECTID      string `json:"PROJECTID"`
	TOPICID        string `json:"TOPICID"`
	SUBSCRIPTIONID string `json:"SUBSCRIPTIONID"`
	TASKTYPE       string `json:"TASKTYPE"`
}

type EnvDB struct {
	USER     string `json:"USER"`
	PASSWORD string `json:"PASSWORD"`
	DBNAME   string `json:"DBNAME"`
	HOST     string `json:"HOST"`
	PORT     string `json:"PORT"`
}

type EnvREDIS struct {
	PASSWORD string `json:"PASSWORD"`
	DB       int    `json:"DB"`
	HOST     string `json:"HOST"`
	PORT     string `json:"PORT"`
}

type EnvGAME struct {
	NO       int    `json:"NO"`
	NAME     string `json:"NAME"`
	CURRENCY string `json:"CURRENCY"`
	PLATFORM string `json:"PLATFORM"`
	AGENT    string `json:"AGENT"`
	HOUSE    string `json:"HOUSE"`
	GROUP    string `json:"GROUP"`
	ABOID    string `json:"ABOID"`
	URL      string `json:"URL"`
	PASSWORD string `json:"PASSWORD"`
	GAMEURL  string `json:"GAMEURL"`
	TABLEKEY string `json:"TABLEKEY"`
	GAMEID   string `json:"GAMEID"`
}

type APIServices struct {
	GamePlayerStation struct {
		Domain   string `json:"Domain"`
		Withdraw string `json:"Withdraw"`
		Deposit  string `json:"Deposit"`
	} `json:"GamePlayerStation"`
}

type EnvONEALL struct {
	PUBLICKEY  string `json:"PUBLICKEY"`
	PRIVATEKEY string `json:"PRIVATEKEY"`
	ENDPOINT   string `json:"ENDPOINT"`
}

type EnvTwilio struct {
	ACCOUNTSID string `json:"ACCOUNTSID"`
	AUTHTOKEN  string `json:"AUTHTOKEN"`
	VERIFYSID  string `json:"VERIFYSID"`
}

type EnvWEBOWNER struct {
	WEBOWNERID string `json:"WEBOWNERID"`
}

type EnvCRON struct {
	CRON_TIME   string `json:"CRON_TIME"`
	TIME_LAYOUT string `json:"TIME_LAYOUT"`
}
