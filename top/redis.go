package top

import (
	"fmt"
	"github.com/go-redis/redis"
)

var (
	RedisClient *redis.Client
)

// NewRedisClient 初始化连接
func NewRedisClient(host string, port, db int, password string, opts ...*redis.Options) (client *redis.Client, err error) {
	opt := new(redis.Options)
	if len(opts) != 0 {
		opt = opts[0]
		if host != "" && port != 0 {
			opt.Addr = fmt.Sprintf("%s:%d", host, port)
		}
		if db != 0 {
			opt.DB = db
		}
		if password != "" {
			opt.Password = password
		}
	} else {
		opt.Addr = fmt.Sprintf("%s:%d", host, port)
		opt.DB = db
		opt.Password = password
	}

	client = redis.NewClient(opt)
	_, err = client.Ping().Result()
	return
}

func GlobalRedisClient(host string, port, db int, password string, opts ...*redis.Options) (err error) {
	client, err := NewRedisClient(host, port, db, password, opts...)
	if err != nil {
		return
	}
	RedisClient = client
	return nil
}

func TestRedisClient() (err error) {
	return GlobalRedisClient("localhost", 6379, 0, "")
}
