package top

import "github.com/hashicorp/consul/api"

var ConsulClient *api.Client

func NewConsulClient(addr string, conf ...*api.Config) (*api.Client, error) {
	if len(conf) == 0 {
		return api.NewClient(&api.Config{Address: addr})
	}
	if addr != "" {
		conf[0].Address = addr
	}
	return api.NewClient(conf[0])
}

func GlobalConsulClient(addr string, conf ...*api.Config) (err error) {
	client, err := NewConsulClient(addr, conf...)
	if err != nil {
		return
	}
	ConsulClient = client
	return nil
}

func TestConsulClient() (err error) {
	ConsulClient, err = api.NewClient(api.DefaultConfig())
	return
}
