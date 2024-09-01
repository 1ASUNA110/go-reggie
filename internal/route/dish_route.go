package route

import (
	"github.com/gin-gonic/gin"
	"go-reggie/internal/api"
)

func SetupDishRoutes(router *gin.Engine) {

	dishApi := api.NewDishApi()

	dishRoutes := router.Group("/dish")
	{
		dishRoutes.GET("/page", dishApi.DishPage)

		dishRoutes.POST("/status/:status", dishApi.DishUpdateStatus)

		dishRoutes.DELETE("/", dishApi.DishDelete)

		dishRoutes.POST("/", dishApi.DishSave)

		dishRoutes.GET("/:id", dishApi.DishGetById)

		dishRoutes.GET("/list", dishApi.DishList)

	}

}
