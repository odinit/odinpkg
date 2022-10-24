package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

var (
	Client *redis.Client
	Nil    = redis.Nil
)

// Init 初始化连接
func Init(host string, port, db int, password string, opts ...*redis.Options) (client *redis.Client, err error) {
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

func DefaultInit() (client *redis.Client, err error) {
	return Init("localhost", 6379, 0, "")
}

func ReplaceGlobal(client *redis.Client) {
	Client = client
}

func GlobalInit(host string, port int, db int, password string, opts ...*redis.Options) (err error) {
	Client, err = Init(host, port, db, password, opts...)
	return
}
