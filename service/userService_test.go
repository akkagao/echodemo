package service

import (
	"echodemo/model"
	"github.com/donnie4w/go-logger/logger"
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

/**
单元测试
测试用户 增、删、改、查操作
*/
func TestUser(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var id int64
	var account string
	Convey("Test UserService", t, func() {
		Convey("InserUser", func() {
			user := model.User{}
			user.Account = "CrazyWolf" + strconv.Itoa(r.Intn(10000))
			user.Username = "CrazyWolf"
			user.Password = "test123"
			user.PhoneNum = "110"
			now := time.Now()
			user.CreateTime = &now

			userid, err := InserUser(user)
			id = userid
			account= user.Account

			So(err, ShouldBeNil)
			So(id, ShouldBeGreaterThan, 1)
		})

		Convey("GetUserById", func() {
			user, err := GetUserById(id)

			logger.Debug(user)
			So(err, ShouldBeNil)
			So(user.Id, ShouldEqual, id)
		})

		Convey("GetUserByAccount", func() {
			user, err := GetUserByAccount(account)

			logger.Debug(user)
			So(err, ShouldBeNil)
			So(user.Id, ShouldEqual, id)
		})

		Convey("UpdateUserPassword", func() {
			password := "CrazyWolf" + strconv.Itoa(r.Intn(100))

			user:=model.User{}
			user.Id = id
			user.Password = password

			count, err := UpdateUserPassword(&user)
			So(err, ShouldBeNil)
			So(count, ShouldEqual, 1)
		})

		Convey("DeleteUser", func() {
			count, err := DeleteUser(id)
			So(err, ShouldBeNil)
			So(count, ShouldEqual, 1)
		})
	})
}
