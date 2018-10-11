package redis_helper

import (
	"time"

	redis "gopkg.in/redis.v4"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(addr, pwd string) *RedisClient {
	option := redis.Options{
		Addr: addr,
	}
	if pwd != "" {
		option.Password = pwd
	}
	client := redis.NewClient(&option)
	return &RedisClient{
		client: client,
	}
}

func (r *RedisClient) Get(key string) (string, error) {
	strCmd := r.client.Get(key)
	return strCmd.Val(), strCmd.Err()
}

func (r *RedisClient) Set(key string, value string,duration time.Duration) {
	r.client.Set(key, value,duration)
}
