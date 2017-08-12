package controller

import (
	"net/http"

	"echodemo/model"
	"echodemo/service"
	"github.com/donnie4w/go-logger/logger"
	"github.com/labstack/echo"
)

/**
注册用户信息
*/
func InserUser(c echo.Context) error {
	userName := c.QueryParam("userName")
	account := c.QueryParam("account")
	password := c.QueryParam("password")
	phoneNum := c.QueryParam("phoneNum")

	user := model.User{}
	user.Username = userName
	user.Password = password
	user.PhoneNum = phoneNum
	user.Account = account

	id, err := service.InserUser(user)

	resultMap := make(map[string]interface{})
	if err != nil {
		logger.Error(err)
		return c.JSON(http.StatusOK, failResut(resultMap))
	}
	resultMap["id"] = id
	return c.JSON(http.StatusOK, successResut(resultMap))
}

/**
修改用户密码
*/
func UpdatePassword(c echo.Context) error {

	account := c.QueryParam("account")
	password := c.QueryParam("password")

	resultMap := make(map[string]interface{})

	user, err := service.GetUserByAccount(account)
	if err != nil {
		logger.Error(err)
		resultMap["errMsg"] = "account not exist"
		return c.JSON(http.StatusOK, failResut(resultMap))
	}

	user.Password = password

	_, err = service.UpdateUserPassword(user)

	if err != nil {
		logger.Error(err)
		resultMap["errMsg"] = "update fail"
		return c.JSON(http.StatusOK, failResut(resultMap))
	}

	return c.JSON(http.StatusOK, successResut(resultMap))
}

/**
查询用户
*/
func SelectUser(c echo.Context) error {

	account := c.QueryParam("account")

	resultMap := make(map[string]interface{})

	user, err := service.GetUserByAccount(account)
	if err != nil || user == nil {
		logger.Error(err)
		resultMap["errMsg"] = "account not exist"
		return c.JSON(http.StatusOK, failResut(resultMap))
	}

	resultMap["data"] = user

	return c.JSON(http.StatusOK, successResut(resultMap))
}

/**
删除用户
*/
func DeleteUser(c echo.Context) error {

	account := c.QueryParam("account")
	resultMap := make(map[string]interface{})

	user, err := service.GetUserByAccount(account)
	if err != nil {
		logger.Error(err)
		resultMap["errMsg"] = "account not exist"
		return c.JSON(http.StatusOK, failResut(resultMap))
	}

	count, err := service.DeleteUser(user.Id)
	if err != nil || count < 1 {
		logger.Error(err)
		resultMap["errMsg"] = "delete fail"
		return c.JSON(http.StatusOK, failResut(resultMap))
	}

	return c.JSON(http.StatusOK, successResut(resultMap))
}


