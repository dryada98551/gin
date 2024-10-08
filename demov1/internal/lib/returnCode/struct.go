package libreturncode

type ReturnCode struct {
	Success     Success     `json:"Success"`
	ServiceFail ServiceFail `json:"ServiceFail"`
	ConnectFail ConnectFail `json:"ConnectFail"`
	GameFail    GameFail    `json:"GameFail"`
}

type payload struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
}

type Success struct {
	Api          payload `json:"Api"`
	Jwt          payload `json:"Jwt"`
	RedisSet     payload `json:"RedisSet"`
	RedisUpdate  payload `json:"RedisUpdate"`
	RedisGet     payload `json:"RedisGet"`
	RedisDelete  payload `json:"RedisDelete"`
	RedisLPush   payload `json:"RedisLPush"`
	RedisRPush   payload `json:"RedisRPush"`
	RedisLPop    payload `json:"RedisLPop"`
	RedisRPop    payload `json:"RedisRPop"`
	RedisZapp    payload `json:"RedisZapp"`
	RedisZRange  payload `json:"RedisZRange"`
	RedisHSet    payload `json:"RedisHSet"`
	RedisHGet    payload `json:"RedisHGet"`
	RedisHGetAll payload `json:"RedisHGetAll"`
	DBInser      payload `json:"DBInser"`
	DBUpdate     payload `json:"DBUpdate"`
	DBSelect     payload `json:"DBSelect"`
	DBDelete     payload `json:"DBDelete"`
	DataParse    payload `json:"DataParse"`
	SocialLogin  payload `json:"SocialLogin"`
	SocialLogout payload `json:"SocialLogout"`
	Publish      payload `json:"Publish"`
	SubReceive   payload `json:"SubReceive"`
	OTPLogin     payload `json:"OTPLogin"`
	OTPLogout    payload `json:"OTPLogout"`
	OTPVerify    payload `json:"OTPVerify"`
	Other        payload `json:"Other"`
}

type ServiceFail struct {
	ApiReq        payload `json:"ApiReq"`
	ApiRes        payload `json:"ApiRes"`
	ApiHeader     payload `json:"ApiHeader"`
	ApiReqCreate  payload `json:"ApiReqCreate"`
	ApiReqDo      payload `json:"ApiReqDo"`
	ApiResRead    payload `json:"ApiResRead"`
	JwtInput      payload `json:"JwtInput"`
	JwtOutput     payload `json:"JwtOutput"`
	RedisKey      payload `json:"RedisKey"`
	RedisInput    payload `json:"RedisInput"`
	RedisOutput   payload `json:"RedisOutput"`
	RedisGet      payload `json:"RedisGet"`
	RedisSet      payload `json:"RedisSet"`
	RedisLPush    payload `json:"RedisLPush"`
	RedisRPush    payload `json:"RedisRPush"`
	RedisLPop     payload `json:"RedisLPop"`
	RedisRPop     payload `json:"RedisRPop"`
	RedisZapp     payload `json:"RedisZapp"`
	RedisZRange   payload `json:"RedisZRange"`
	RedisHSet     payload `json:"RedisHSet"`
	RedisHGet     payload `json:"RedisHGet"`
	RedisHGetAll  payload `json:"RedisHGetAll"`
	DBTable       payload `json:"DBTable"`
	DBData        payload `json:"DBData"`
	DBInser       payload `json:"DBInser"`
	DBUpdate      payload `json:"DBUpdate"`
	DBInput       payload `json:"DBInput"`
	DBOutput      payload `json:"DBOutput"`
	DataParse     payload `json:"DataParse"`
	DataUndefined payload `json:"DataUndefined"`
	DataToJson    payload `json:"DataToJson"`
	Publish       payload `json:"Publish"`
	SubReceive    payload `json:"SubReceive"`
	SocialReq     payload `json:"SocialReq"`
	SocialSet     payload `json:"SocialSet"`
	SocialApi     payload `json:"SocialApi"`
	SocialRes     payload `json:"SocialRes"`
	Other         payload `json:"Other"`
}

type ConnectFail struct {
	Api   payload `json:"Api"`
	Jwt   payload `json:"Jwt"`
	Redis payload `json:"Redis"`
	DB    payload `json:"DB"`
	Other payload `json:"Other"`
}

type GameFail struct {
	GameNo            payload `json:"GameNo"`
	SysCreateAccount  payload `json:"SysCreateAccount"`
	GameCreateAccount payload `json:"GameCreateAccount"`
	GameAccountExist  payload `json:"GameAccountExist"`
	Other             payload `json:"Other"`
}
