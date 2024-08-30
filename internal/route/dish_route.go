package route

import (
	"github.com/gin-gonic/gin"
	"go-reggie/internal/api"
)

func SetupDishRoutes(router *gin.Engine) {

	dishApi := api.NewDishApi()

	categoryRoutes := router.Group("/dish")
	{
		categoryRoutes.GET("/page", dishApi.DishPage)

		categoryRoutes.POST("/status/:status", dishApi.DishUpdateStatus)

		categoryRoutes.DELETE("/", dishApi.DishDelete)
	}

}
