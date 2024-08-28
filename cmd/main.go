package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-reggie/config"
	"go-reggie/internal/bootstrap"
	"go-reggie/internal/global"
	"go-reggie/internal/middleware"
	"go-reggie/internal/route"
)

func Start() {
	config.InitConfig()

}

func Clean() {
	fmt.Println("============Clean============")
}

func main() {
	defer Clean()
	Start()

	// 创建 Gin 引擎
	r := gin.Default()

	// 强制设置日志颜色
	gin.ForceConsoleColor()

	// 提供静态文件服务
	r.Static("/", "./static")

	// 使用cookie存储session
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("reggieSession", store))

	// 权限校验中间件
	r.Use(middleware.AuthMiddleware())

	//

	//从内部包中初始化数据库
	db, err := bootstrap.InitDB()
	if err != nil {
		panic(err)
	}

	global.DB = db

	route.SetupRouter(r)

	// 启动 HTTP 服务器
	r.Run(":" + viper.GetString("sever.port"))

}
