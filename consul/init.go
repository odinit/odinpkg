package consul

import "github.com/hashicorp/consul/api"

var Client *api.Client
var KV *api.KV
var Agent *api.Agent
var Health *api.Health

//var

var QueryOptions *api.QueryOptions
var WriteOptions *api.WriteOptions

func NewClient(conf *api.Config) (*api.Client, error) {
	return api.NewClient(conf)
}

func NewClientWithDefaultConfig() (*api.Client, error) {
	return api.NewClient(api.DefaultConfig())
}

func GlobalClient(conf *api.Config) error {
	client, err := api.NewClient(conf)
	if err != nil {
		return err
	}
	ReplaceGlobal(client)
	return nil
}

func GlobalClientWithDefaultConfig() error {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return err
	}
	ReplaceGlobal(client)
	return nil
}

func ReplaceGlobal(client *api.Client) {
	Client = client
	Agent = Client.Agent()
	KV = Client.KV()
	Health = Client.Health()

}
