package filter

import (
	"echodemo/model"
	"github.com/labstack/echo"
	"net/http"
	"echodemo/service"
)

var (
	User *model.User
)
/**
此方法为公共拦截器
这里做是否登录、权限校验等操作
给全局变量设值可以在Controller 中获取到
 */
func OpenTracing() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 从请求头中获取请求参数
			//token := c.Request().Header.Get("t")
			//uid := c.Request().Header.Get("u")
			//id,err := strconv.ParseInt(uid, 10, 64)
			//if err != nil {
			//	return c.JSON(http.StatusUnauthorized, "error request")
			//}
			user, err := service.GetUserById(1)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "error request")
			}
			User = user
			//logger.Debug("token:", token, "uid:", uid)
			return next(c)
		}
	}
}
