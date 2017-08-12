package main

import (
	"echodemo/conf"
	"echodemo/router"
	"echodemo/service"
	"github.com/donnie4w/go-logger/logger"
	"os"
)

func main() {

	//如果没有设置启动参数默认为dev模式
	var runmodel = "dev"

	//读取启动参数，作为运行模式，选择不同的配置文件
	if len(os.Args) == 2 {
		runmodel = os.Args[1]
		logger.Debug("run model:", runmodel)
	}

	//读取配置文件并且初始化配置对象
	conf.InitConf(runmodel)

	// 初始化日志组件
	initLog()

	//初始化数据库连接
	service.Start()

	//启动echo
	router.Start()
}

/**
 * 初始化日志设置
 */
func initLog() {

	logger.SetConsole(conf.Config.Log.PrintConsole)

	//指定日志文件备份方式为文件大小的方式
	//第一个参数为日志文件存放目录
	//第二个参数为日志文件命名
	//第三个参数为备份文件最大数量
	//第四个参数为备份文件大小
	//第五个参数为文件大小的单位
	logger.SetRollingFile(conf.Config.Log.LogPath, conf.Config.Log.LogFileName, conf.Config.Log.MaxNumber, conf.Config.Log.MaxSize, logger.MB)

	//指定日志文件备份方式为日期的方式
	//第一个参数为日志文件存放目录
	//第二个参数为日志文件命名
	// logger.SetRollingDaily("/Users/crazywolf/tmp/log", "test.log")

	//指定日志级别  ALL，DEBUG，INFO，WARN，ERROR，FATAL，OFF 级别由低到高
	//一般习惯是测试阶段为debug，生成环境为info以上
	switch conf.Config.Log.Level {
	case 0:
		logger.SetLevel(logger.ALL)
	case 1:
		logger.SetLevel(logger.DEBUG)
	case 2:
		logger.SetLevel(logger.INFO)
	case 3:
		logger.SetLevel(logger.WARN)
	case 4:
		logger.SetLevel(logger.ERROR)
	case 5:
		logger.SetLevel(logger.FATAL)
	case 6:
		logger.SetLevel(logger.OFF)
	default:
		logger.SetLevel(logger.DEBUG)
	}
}
