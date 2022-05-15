package gredis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ali-sharafi/wallet/pkg/settings"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func Setup() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     settings.RedisSetting.Host,
		Password: settings.RedisSetting.Password,
		DB:       0,
	})

	return rdb
}

func Set(key string, data interface{}, expireTime time.Duration) error {
	rdb := Setup()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = rdb.Set(ctx, key, value, expireTime*time.Second).Err()

	if err != nil {
		return err
	}

	return nil
}

func Get(key string) (interface{}, error) {
	rdb := Setup()
	var result interface{}
	value, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("%v does not exists", key)
	} else if err != nil {
		return nil, err
	} else {
		json.Unmarshal([]byte(value), &result)
		return result, nil
	}
}
