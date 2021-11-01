package main

import (
	"ea-competition-api/boot/db/mysql"
	"ea-competition-api/boot/log"
	"ea-competition-api/config"
	token2 "ea-competition-api/library/token"
	"ea-competition-api/router"
	"fmt"

	"go.uber.org/zap"
)

func main() {

	// 初始化配置文件
	config.LoadConfig()

	//初始化log
	log.Init("logs")

	//mysql初始化
	err := mysql.Init(config.Conf().MySQL.Dsn)
	if err != nil {
		log.Logger().Error(" mysql connect error", zap.Error(err))
		return
	}

	// redis
	//redis.Connect()

	//es.Connect()

	// nsq consumer
	//nsq.NewNsq().InitConsumer("create_bit_topic", "create_bit_channel", &participate.CreateBitConsumer{})

	//初始化认证token
	token := new(token2.Token)
	token.InitToken(config.Conf().API.AuthToken)
	//gin路由引擎配置
	engine := router.InitRouter(token, log.Logger())
	//启动服务
	if err := engine.Run(fmt.Sprintf("%s:%d", config.Conf().Application.Host, config.Conf().Application.Port)); err != nil {
		log.Logger().Error("start service error", zap.Error(err))
	}
}
