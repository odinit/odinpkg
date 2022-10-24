package top

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

var ElasticClient *elastic.Client

// NewElasticClient 初始化
func NewElasticClient(host string, port int, proto string) (client *elastic.Client, err error) {
	url_ := fmt.Sprintf("%s://%s:%d", proto, host, port)
	client, err = elastic.NewClient(
		elastic.SetSniff(false), // 禁用嗅探:不加上elastic.SetSniff(false) 会连接不上
		elastic.SetURL(url_))
	if err != nil {
		return
	}
	_, _, err = client.Ping(url_).Do(context.Background())
	if err != nil {
		return
	}
	_, err = client.ElasticsearchVersion(url_)
	if err != nil {
		return
	}
	return
}

func GlobalElasticClient(host string, port int, proto string) (err error) {
	client, err := NewElasticClient(host, port, proto)
	if err != nil {
		return
	}
	ElasticClient = client
	return nil
}
