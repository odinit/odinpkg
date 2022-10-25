package top

import (
	"fmt"
	"github.com/go-redis/redis"
)

var (
	RedisClient *redis.Client
)

// NewRedisClient 初始化连接
func NewRedisClient(host string, port, db int, password string, opts *redis.Options) (client *redis.Client, err error) {
	if opts == nil {
		opts = &redis.Options{}
	}
	opts.Addr = fmt.Sprintf("%s:%d", host, port)
	opts.Password = password
	opts.DB = db

	client = redis.NewClient(opts)
	_, err = client.Ping().Result()
	return
}

func GlobalRedisClient(host string, port, db int, password string, opts *redis.Options) (err error) {
	client, err := NewRedisClient(host, port, db, password, opts)
	if err != nil {
		return
	}
	RedisClient = client
	return nil
}

func TestRedisClient() (err error) {
	return GlobalRedisClient("localhost", 6379, 0, "", nil)
}
