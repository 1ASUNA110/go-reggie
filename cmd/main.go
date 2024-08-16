package main

import "github.com/gin-gonic/gin"

func main() {
	// 创建 Gin 引擎
	r := gin.Default()

	// 强制设置日志颜色
	gin.ForceConsoleColor()

	// 定义路由和处理函数
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 启动 HTTP 服务器
	r.Run() // 默认监听 :8080
}
