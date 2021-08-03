package public

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var redisAdapter *RedisAdapter

func init()  {
	 redisAdapter = &RedisAdapter{
	 	cli: redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})}
}

type RedisCallBackFn func() (interface{}, error)

func GetRedisAdapter() *RedisAdapter {
	return redisAdapter
}

type RedisAdapter struct {
	cli *redis.Client
}



func (r *RedisAdapter) GetValue(fn RedisCallBackFn, key string) (interface{}, error) {
	var ctx = context.Background()
	res, err := r.cli.Get(ctx, key).Result()
	if err == redis.Nil {
		return fn()
	} else if err != nil {
		log.Error().Err(err)
		return fn()
	} else {
		return res, err
	}
}

func (r *RedisAdapter) UpdateValue(fn RedisCallBackFn, key string) (interface{}, error) {
	res, err := fn()
	if err != nil {
		return res, err
	}
	var  ctx = context.Background()
	_, err = r.cli.Del(ctx, key).Result()
	if err == redis.Nil {
		return res, err
	} else if err != nil {
		log.Error().Err(err)
		return res, err
	} else {
		return res, err
	}
}



func (r *RedisAdapter) DeleteValue(fn RedisCallBackFn,key string) (interface{}, error)  {
	res, err := fn()
	if err != nil {
		return res, err
	}
	var ctx = context.Background()
	_, err = r.cli.Del(ctx, key).Result()
	if err == redis.Nil {
		return res, err
	} else if err != nil {
		log.Error().Err(err)
		return res, err
	} else {
		return res, err
	}
}