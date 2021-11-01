package es

import (
	"context"
	"ea-competition-api/boot/log"
	"ea-competition-api/config"
	"sync"

	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

var (
	err    error
	once   sync.Once
	client *elastic.Client
)

//集群连接
func Connect() {
	once.Do(func() {
		urls := config.Conf().Es.Hosts
		client, err = elastic.NewClient(elastic.SetURL(urls...), elastic.SetBasicAuth(config.Conf().Es.Username, config.Conf().Es.Password), elastic.SetSniff(false))
		if err != nil {
			panic("Elasticsearch connect err：" + err.Error())
		}

		// ping ip 检查
		for _, url := range urls {
			if _, _, err := client.Ping(url).Do(context.Background()); err != nil {
				log.Logger().Error("Elasticsearch ping ip err: "+url, zap.Error(err))
				panic("Elasticsearch ping ip err" + err.Error())
			}
		}
	})
}

// ES服务
func Client() *elastic.Client {
	return client
}
