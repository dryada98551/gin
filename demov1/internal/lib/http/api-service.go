package libhttp

type HttpUrl struct {
	GameLoginAPI   payload `json:"game-login-api"`
	GamePlayerAPI  payload `json:"game-player-api"`
	ChargeAPI      payload `json:"charge-api"`
	KYCAPI         payload `json:"kyc-api"`
	LoginAPI       payload `json:"login-api"`
	MemberLevelAPI payload `json:"member-level-api"`
	WithdrawAPI    payload `json:"withdraw-api"`
}
type payload struct {
	Http  string `json:"Http"`
	// Https string `json:"Https"`
}

var HttpUrlDebug = HttpUrl{
	GameLoginAPI: payload{
		Http:  "http://bab-02-game-login-api:8080",
	},
	GamePlayerAPI: payload{
		Http:  "http://bab-04-game-player-api:8080",
	},
	ChargeAPI: payload{
		Http:  "http://bab-06-charge-api:8080",
	},
	KYCAPI: payload{
		Http:  "http://bab-07-kyc-api:8080",
	},
	LoginAPI: payload{
		Http:  "http://bab-08-login-api:8080",
	},
	MemberLevelAPI: payload{
		Http:  "http://bab-09-member-level-api:8080",
	},
	WithdrawAPI: payload{
		Http:  "http://bab-11-withdraw-api:8080",
	},
}

var HttpUrlUat = HttpUrl{
	GameLoginAPI: payload{
		Http:  "",
	},
	GamePlayerAPI: payload{
		Http:  "",
	},
	ChargeAPI: payload{
		Http:  "",
	},
	KYCAPI: payload{
		Http:  "",
	},
	LoginAPI: payload{
		Http:  "",
	},
	MemberLevelAPI: payload{
		Http:  "",
	},
	WithdrawAPI: payload{
		Http:  "",
	},
}

var HttpUrlRelease = HttpUrl{
	GameLoginAPI: payload{
		Http:  "",
	},
	GamePlayerAPI: payload{
		Http:  "",
	},
	ChargeAPI: payload{
		Http:  "",
	},
	KYCAPI: payload{
		Http:  "",
	},
	LoginAPI: payload{
		Http:  "",
	},
	MemberLevelAPI: payload{
		Http:  "",
	},
	WithdrawAPI: payload{
		Http:  "",
	},
}
