package lib

import (
	"context"
	libmodel "main/internal/lib/model"
	libredis "main/internal/lib/redis"
	libreturncode "main/internal/lib/returnCode"
)

func LoginCheck(ctx context.Context, JwtRedisKey string, UserID string) libmodel.SystemReturn {
	// Create a context

	key := JwtRedisKey + ":" + UserID

	_, sysReturn := libredis.RedisGet(ctx, key)
	if sysReturn.Code != libreturncode.Cobelib.Success.Other.Code {
		return sysReturn
	}

	return sysReturn
}
