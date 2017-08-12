package util

import "time"

import "strconv"

/**
获取系统当前时间毫秒数
*/
func Millisecond() int64 {
	return time.Now().UnixNano() / 1000000
}

func FormatTime(date int64) string {
	ago := (Millisecond() - date) / 1000
	if ago < 60*10 {
		return "刚刚"
	} else if (ago < 60*60) && (ago >= 60*10) {
		s := ago / 60
		return strconv.FormatInt(s, 10) + "分钟前"
	} else if (ago >= 60*60) && (ago < 60*60*24) {
		s := ago / (60 * 60)
		return strconv.FormatInt(s, 10) + "小时前"
	} else if (ago >= 60*60*24) && (ago < 60*60*24*30) {
		s := ago / (60 * 60 * 24)
		return strconv.FormatInt(s, 10) + "天前"
	} else if (ago >= 60*60*24*30) && (ago < 60*60*24*30*12) {
		s := ago / (60 * 60 * 24 * 30)
		return strconv.FormatInt(s, 10) + "月前"
	} else if ago >= 60*60*24*30*12 {
		s := ago / (60 * 60 * 24 * 30 * 12)
		return strconv.FormatInt(s, 10) + "年前"
	}
	return "很久以前"
}
