package middleware

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-reggie/internal/model/vo/response"
	"regexp"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1、获取本次请求的URI
		requestURI := c.Request.RequestURI
		session := sessions.Default(c)
		employeeID := session.Get("employee")

		// 输出URI日志
		fmt.Println("Request URI:", requestURI)

		allowedUrls := []string{
			"^/employee/login$",
			"^/employee/logout$",
			"^/static/backend/.*$",
			"^/static/front/.*$",
		}

		// 2、判断本次请求是否需要登录认证，如果不需要，直接放行
		for _, pattern := range allowedUrls {
			matched, _ := regexp.MatchString(pattern, requestURI)
			if matched {
				c.Next()
				return
			}
		}

		// 3、判断登录状态，如果已登录，则直接放行

		if employeeID != nil {
			c.Next()
			return
		}

		// 4、如果未登录,则返回未登录结果
		response.Fail(response.LOGIN_CHECK_ERROR(), c)
		c.Abort()

	}
}
