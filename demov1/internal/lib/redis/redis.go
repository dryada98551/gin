package libredis

import (
	"context"

	"github.com/go-redis/redis/v8"

	configs "main/internal/lib/init"
	libmodel "main/internal/lib/model"
	libreturncode "main/internal/lib/returnCode"
)

func RedisGet(ctx context.Context, key string) (string, libmodel.SystemReturn) {
	// Create a context
	var sysReturn libmodel.SystemReturn

	value, err := configs.Global.Redis.Client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			sysReturn.Code = libreturncode.Cobelib.ServiceFail.RedisKey.Code
			sysReturn.Message = libreturncode.Cobelib.ServiceFail.RedisKey.Message
			return value, sysReturn
		} else {
			sysReturn.Code = libreturncode.Cobelib.ServiceFail.RedisGet.Code
			sysReturn.Message = libreturncode.Cobelib.ServiceFail.RedisGet.Message
			return value, sysReturn
		}
	}

	sysReturn.Code = libreturncode.Cobelib.Success.Other.Code
	sysReturn.Message = libreturncode.Cobelib.Success.RedisGet.Message
	return value, sysReturn
}

func RedisSet(ctx context.Context, key string, data string) (libmodel.SystemReturn) {
	// Create a context
	var sysReturn libmodel.SystemReturn

	err := configs.Global.Redis.Client.Set(ctx, key, data, 0).Err()
	if err != nil {
		sysReturn.Code = libreturncode.Cobelib.ServiceFail.RedisSet.Code
		sysReturn.Message = libreturncode.Cobelib.ServiceFail.RedisSet.Message
		return sysReturn
	}

	sysReturn.Code = libreturncode.Cobelib.Success.Other.Code
	sysReturn.Message = libreturncode.Cobelib.Success.RedisSet.Message
	return sysReturn
}

func RedisLPush(ctx context.Context, key string, data string) (libmodel.SystemReturn) {
	// Create a context
	var sysReturn libmodel.SystemReturn

	err := configs.Global.Redis.Client.LPush(ctx, key, data).Err()
	if err != nil {
		sysReturn.Code = libreturncode.Cobelib.ServiceFail.RedisLPush.Code
		sysReturn.Message = libreturncode.Cobelib.ServiceFail.RedisLPush.Message
		return sysReturn
	}

	sysReturn.Code = libreturncode.Cobelib.Success.Other.Code
	sysReturn.Message = libreturncode.Cobelib.Success.RedisSet.Message
	return sysReturn
}

func RedisRPush(ctx context.Context, key string, data string) (libmodel.SystemReturn) {
	// Create a context
	var sysReturn libmodel.SystemReturn

	err := configs.Global.Redis.Client.RPush(ctx, key, data).Err()
	if err != nil {
		sysReturn.Code = libreturncode.Cobelib.ServiceFail.RedisRPush.Code
		sysReturn.Message = libreturncode.Cobelib.ServiceFail.RedisRPush.Message
		return sysReturn
	}

	sysReturn.Code = libreturncode.Cobelib.Success.Other.Code
	sysReturn.Message = libreturncode.Cobelib.Success.RedisSet.Message
	return sysReturn
}

func RedisLPop(ctx context.Context, key string) (string, libmodel.SystemReturn) {
	// Create a context
	var sysReturn libmodel.SystemReturn

	val, err := configs.Global.Redis.Client.LPop(ctx, key).Result()
	if err != nil {
		sysReturn.Code = libreturncode.Cobelib.ServiceFail.RedisLPop.Code
		sysReturn.Message = libreturncode.Cobelib.ServiceFail.RedisLPop.Message
		return "", sysReturn
	}

	sysReturn.Code = libreturncode.Cobelib.Success.Other.Code
	sysReturn.Message = libreturncode.Cobelib.Success.RedisSet.Message
	return val, sysReturn
}

func RedisRPop(ctx context.Context, key string) (string, libmodel.SystemReturn) {
	// Create a context
	var sysReturn libmodel.SystemReturn

	val, err := configs.Global.Redis.Client.RPop(ctx, key).Result()
	if err != nil {
		sysReturn.Code = libreturncode.Cobelib.ServiceFail.RedisRPop.Code
		sysReturn.Message = libreturncode.Cobelib.ServiceFail.RedisRPop.Message
		return "", sysReturn
	}

	sysReturn.Code = libreturncode.Cobelib.Success.Other.Code
	sysReturn.Message = libreturncode.Cobelib.Success.RedisSet.Message
	return val, sysReturn
}

func RedisZAdd(ctx context.Context, key string, data *redis.Z) (libmodel.SystemReturn) {
	// Create a context
	var sysReturn libmodel.SystemReturn

	_, err := configs.Global.Redis.Client.ZAdd(ctx, key, data).Result()
	if err != nil {
		sysReturn.Code = libreturncode.Cobelib.ServiceFail.RedisZapp.Code
		sysReturn.Message = libreturncode.Cobelib.ServiceFail.RedisZapp.Message
		return sysReturn
	}

	sysReturn.Code = libreturncode.Cobelib.Success.Other.Code
	sysReturn.Message = libreturncode.Cobelib.Success.RedisSet.Message
	return sysReturn
}

func RedisZRange(ctx context.Context, key string) ([]string, libmodel.SystemReturn) {
	// Create a context
	var sysReturn libmodel.SystemReturn

	val, err := configs.Global.Redis.Client.ZRange(ctx, key, 0, -1).Result()
	if err != nil {
		sysReturn.Code = libreturncode.Cobelib.ServiceFail.RedisZRange.Code
		sysReturn.Message = libreturncode.Cobelib.ServiceFail.RedisZRange.Message
		return nil, sysReturn
	}

	sysReturn.Code = libreturncode.Cobelib.Success.Other.Code
	sysReturn.Message = libreturncode.Cobelib.Success.RedisSet.Message
	return val, sysReturn
}

func RedisHSet(ctx context.Context, key string, item string, value string) (libmodel.SystemReturn) {
	// Create a context
	var sysReturn libmodel.SystemReturn

	_, err := configs.Global.Redis.Client.HSet(ctx, key, item, value).Result()
	if err != nil {
		sysReturn.Code = libreturncode.Cobelib.ServiceFail.RedisHSet.Code
		sysReturn.Message = libreturncode.Cobelib.ServiceFail.RedisHSet.Message
		return sysReturn
	}

	sysReturn.Code = libreturncode.Cobelib.Success.Other.Code
	sysReturn.Message = libreturncode.Cobelib.Success.RedisSet.Message
	return sysReturn
}

func RedisHGet(ctx context.Context, key string, item string) (string, libmodel.SystemReturn) {
	// Create a context
	var sysReturn libmodel.SystemReturn

	val, err := configs.Global.Redis.Client.HGet(ctx, key, item).Result()
	if err != nil {
		sysReturn.Code = libreturncode.Cobelib.ServiceFail.RedisHGet.Code
		sysReturn.Message = libreturncode.Cobelib.ServiceFail.RedisHGet.Message
		return "", sysReturn
	}

	sysReturn.Code = libreturncode.Cobelib.Success.Other.Code
	sysReturn.Message = libreturncode.Cobelib.Success.RedisSet.Message
	return val, sysReturn
}

func RedisHGetAll(ctx context.Context, key string) (map[string]string, libmodel.SystemReturn) {
	// Create a context
	var sysReturn libmodel.SystemReturn

	val, err := configs.Global.Redis.Client.HGetAll(ctx, key).Result()
	if err != nil {
		sysReturn.Code = libreturncode.Cobelib.ServiceFail.RedisHGetAll.Code
		sysReturn.Message = libreturncode.Cobelib.ServiceFail.RedisHGetAll.Message
		return nil, sysReturn
	}

	sysReturn.Code = libreturncode.Cobelib.Success.Other.Code
	sysReturn.Message = libreturncode.Cobelib.Success.RedisSet.Message
	return val, sysReturn
}
