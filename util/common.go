package util

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

func M5dToken(id int64) string {
	h := md5.New()
	h.Write([]byte(strconv.FormatInt(id, 10)))
	tokenByte := h.Sum([]byte("CrazyWolf"))
	return hex.EncodeToString(tokenByte)
}

func M5dTokenStr(id string) string {
	h := md5.New()
	h.Write([]byte(id))
	tokenByte := h.Sum([]byte("CrazyWolf"))
	return hex.EncodeToString(tokenByte)
}

func Join(int64Slice []int64, sep string) string {
	strSlice := make([]string, len(int64Slice), len(int64Slice))
	for i, value := range int64Slice {
		strSlice[i] = strconv.FormatInt(value, 10)
	}
	return strings.Join(strSlice, ",")
}

func GetWenStr(len int) string {
	wenSlice := make([]string, len, len)
	for index := 0; index < len; index++ {
		wenSlice[index] = "?"
	}
	return strings.Join(wenSlice, ",")
}

func SubString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(rs[begin:end])
}
