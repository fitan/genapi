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

func GetRedisAdapter() *RedisAdapter {
	return redisAdapter
}

type RedisAdapter struct {
	cli *redis.Client
}



func (r *RedisAdapter) GetValueByKey(key string, fc func() (interface{}, error)) (interface{},error) {
	var ctx = context.Background()
	res, err  := r.cli.Get(ctx, key).Result()
	if err == redis.Nil {
		return fc()
	} else if err != nil {
		return fc()
	} else {
		return  res, err
	}
}

func (r *RedisAdapter) DeleteValueByKey(key string) error  {
	var ctx = context.Background()
	del := r.cli.Del(ctx, key)
	_, err := del.Result()
	return err
}