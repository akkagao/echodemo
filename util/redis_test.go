package util

import (
//"fmt"
//"github.com/garyburd/redigo/redis"
//"testing"
//
//"github.com/pquerna/ffjson/ffjson"
//. "github.com/smartystreets/goconvey/convey"
//"sync"
)

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

//func TestGetRedisConn(t *testing.T) {
//	Convey("", t, func() {
//
//		user1 := &User{}
//		getObjFormRedis(user1, createuser)
//		fmt.Println(user1)
//	})
//}
//
//func createuser() {
//	user := &User{}
//	user.Id = 1
//	user.Name = "akka"
//	addToRedis(user)
//}

//func addToRedis(obj interface{}) {
//	conn := GetRedisConn()
//	defer conn.Close()
//	value, _ := ffjson.Marshal(obj)
//	conn.Do("SET", "user_0001", value)
//}
//
//func getObjFormRedis(obj interface{}, function func()) {
//	conn := GetRedisConn()
//	defer conn.Close()
//	value, _ := redis.Bytes(conn.Do("GET", "user_0001"))
//	ffjson.Unmarshal(value, obj)
//}
//
//func BenchmarkGetRedisConn(t *testing.B) {
//	var sw sync.WaitGroup
//	for i := 0; i < 100000; i++ {
//		value := i
//		if i%3 == 0 {
//			sw.Add(1)
//			go excute(value)
//			sw.Done()
//			continue
//		}
//		excute(value)
//	}
//	sw.Wait()
//}
//
//func excute(value int) int {
//	conn := GetRedisConn()
//	defer conn.Close()
//	key := fmt.Sprintf("test_%d", value)
//	conn.Do("SET", key, value)
//	result, err := redis.Int(conn.Do("GET", key))
//	if value != result {
//		fmt.Println("Error..", err, " -|- ", err.Error())
//	}
//	return result
//}
