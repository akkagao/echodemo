package service

import (
	"echodemo/conf"
	"github.com/donnie4w/go-logger/logger"
	_ "github.com/go-sql-driver/mysql"
)
/**
主要用于单元测试初始化环境
在此目录下执行 go test  会初始化 日志，读取配置文件、数据库链接等操作
 */

/**
启动日志 数据库链接等依赖
*/
func init() {
	//读取配置文件并且初始化配置对象
	conf.InitTestConf("dev")
	// 初始化日志组件
	initLog()
	Start()
	logger.Debug("测试init")
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
