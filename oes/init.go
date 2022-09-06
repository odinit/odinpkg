package ges

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

var Client *elastic.Client

// Init 初始化
func Init(host string, port int, proto string) (err error) {
	url_ := fmt.Sprintf("%s://%s:%d", proto, host, port)
	Client, err = elastic.NewClient(
		elastic.SetSniff(false), // 禁用嗅探:不加上elastic.SetSniff(false) 会连接不上
		elastic.SetURL(url_))
	if err != nil {
		return
	}
	_, _, err = Client.Ping(url_).Do(context.Background())
	if err != nil {
		return
	}
	_, err = Client.ElasticsearchVersion(url_)
	if err != nil {
		return
	}
	return
}
