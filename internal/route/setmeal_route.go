package route

import (
	"github.com/gin-gonic/gin"
	"go-reggie/internal/api"
)

func SetupSetmealRoutes(router *gin.Engine) {

	setmealApi := api.NewSetmealApi()

	setmealRoutes := router.Group("/setmeal")
	{
		setmealRoutes.POST("/", setmealApi.SetmealSave)

		setmealRoutes.GET("/page", setmealApi.SetmealPage)

		setmealRoutes.POST("/status/:status", setmealApi.SetmealUpdateStatus)

		//setmealRoutes.DELETE("/", setmealApi.SetmealDelete)
		//
		//setmealRoutes.PUT("/", setmealApi.SetmealUpdate)
		//
		//setmealRoutes.GET("/list", setmealApi.SetmealList)
	}

}
