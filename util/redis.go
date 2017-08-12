package util

//import (
//	"time"
//
//	"github.com/astaxie/beego"
//	"github.com/garyburd/redigo/redis"
//)

//var (
//	// 定义常量
//	redisClient *redis.Pool
//)
//
///**
//获取redis链接对象
//demo ：
//	conn := util.GetRedisConn()
//	defer conn.Close()
//*/
//func GetRedisConn() redis.Conn {
//	return redisClient.Get()
//}

/**
初始化redis连接池
*/
func startRedisConn() {
	// 从配置文件获取redis的ip以及db
	//redis_host := "127.0.0.1:6379"
	//redis_host := beego.AppConfig.String("redis_host")
	//redis_db, _ := beego.AppConfig.Int("redis_db")
	//beego.Debug(redis_host)
	//// 建立连接池
	//redisClient = &redis.Pool{
	//	// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
	//	MaxIdle:     beego.AppConfig.DefaultInt("redis.maxidle", 1),
	//	MaxActive:   beego.AppConfig.DefaultInt("redis.maxactive", 10000),
	//	IdleTimeout: 180 * time.Second,
	//	Dial: func() (redis.Conn, error) {
	//		c, err := redis.Dial("tcp", redis_host)
	//		if err != nil {
	//			return nil, err
	//		}
	//		// 选择db
	//		c.Do("SELECT", redis_db)
	//		return c, nil
	//	},
	//}
	//
	////判断redis是否启动成功
	//result, err := GetRedisConn().Do("PING")
	//if err != nil || result != "PONG" {
	//	beego.Error("connect redis fail,please check redis services ! 3 Second leater will retry!")
	//	time.Sleep(3 * time.Second)
	//	startRedisConn()
	//}
}

func init() {
	startRedisConn()
}
