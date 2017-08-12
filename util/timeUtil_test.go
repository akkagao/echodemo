package util

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFormatTime(t *testing.T) {
	Convey("func FormatTime(date int64) string  ", t, func() {
		Convey("刚刚1", func() {
			// 当前毫秒数
			So(FormatTime(time.Now().UnixNano()/1000000), ShouldEqual, "刚刚")
		})
		Convey("刚刚2", func() {
			// 9分钟前
			s, _ := time.ParseDuration("-9m")
			So(FormatTime(time.Now().Add(s).UnixNano()/1000000), ShouldEqual, "刚刚")
		})
		Convey("10分钟前", func() {
			// 10分钟前
			s, _ := time.ParseDuration("-10m")
			So(FormatTime(time.Now().Add(s).UnixNano()/1000000), ShouldEqual, "10分钟前")
		})
		Convey("61分钟前刚刚", func() {
			// 61分钟前
			s, _ := time.ParseDuration("-61m")
			So(FormatTime(time.Now().Add(s).UnixNano()/1000000), ShouldEqual, "1小时前")
		})
		Convey("1小时前", func() {
			// 1小时前
			s, _ := time.ParseDuration("-1h")
			So(FormatTime(time.Now().Add(s).UnixNano()/1000000), ShouldEqual, "1小时前")
		})
		Convey("23小时前", func() {
			// 23小时前
			s, _ := time.ParseDuration("-23h")
			So(FormatTime(time.Now().Add(s).UnixNano()/1000000), ShouldEqual, "23小时前")
		})
		Convey("24小时前", func() {
			// 24小时前
			s, _ := time.ParseDuration("-24h")
			So(FormatTime(time.Now().Add(s).UnixNano()/1000000), ShouldEqual, "1天前")
		})
		Convey("50小时前", func() {
			// 50小时前
			s, _ := time.ParseDuration("-50h")
			So(FormatTime(time.Now().Add(s).UnixNano()/1000000), ShouldEqual, "2天前")
		})
		Convey("720小时前", func() {
			// 720小时前(一个月)
			s, _ := time.ParseDuration("-720h")
			So(FormatTime(time.Now().Add(s).UnixNano()/1000000), ShouldEqual, "1月前")
		})
		Convey("8640小时前", func() {
			// 8640小时前(一年)
			s, _ := time.ParseDuration("-8640h")
			So(FormatTime(time.Now().Add(s).UnixNano()/1000000), ShouldEqual, "1年前")
		})
		Convey("25930小时前", func() {
			// 25930小时前(三年)
			s, _ := time.ParseDuration("-25930h")
			So(FormatTime(time.Now().Add(s).UnixNano()/1000000), ShouldEqual, "3年前")
		})
	})
}
