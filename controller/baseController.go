package controller

import (
	"echodemo/model"
	"github.com/labstack/echo"
)

type BaseController struct {
	echo.Context
	jsonData map[string]interface{}
	userme   model.User
}

/*
组装执行成功返回消息
可以设置公共参数
 */
func successResut(resutMap map[string]interface{}) map[string]interface{} {
	resutMap["state"] = "success"
	return resutMap
}

/*
组装执行失败返回消息
可以设置公共参数
 */
func failResut(resutMap map[string]interface{}) map[string]interface{} {
	resutMap["state"] = "fail"
	return resutMap
}
