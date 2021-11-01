package main

import (
	"ea-competition-api/boot/db/mysql"
	"ea-competition-api/boot/log"
	"ea-competition-api/config"
	"ea-competition-api/job/apply"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func main() {

	// todo 后期这些初始化 统一封装 和 业务项目共用
	// 初始化配置文件
	config.LoadConfig()

	//初始化log
	log.Init("logs/job")

	//mysql初始化
	err := mysql.Init(config.Conf().MySQL.Dsn)
	if err != nil {
		log.Logger().Error(" mysql connect error", zap.Error(err))
		return
	}

	initCron()
}

func initCron() {

	c := cron.New(cron.WithSeconds())

	_, _ = c.AddFunc("*/5 * * * * *", apply.CheckStatus)

	c.Start()

	select {}
}
