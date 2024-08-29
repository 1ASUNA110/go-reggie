package route

import (
	"github.com/gin-gonic/gin"
	"go-reggie/internal/api"
)

func SetupCategoryRoutes(router *gin.Engine) {

	categoryApi := api.NewCategoryApi()

	categoryRoutes := router.Group("/category")
	{
		categoryRoutes.POST("/", categoryApi.CategorySave)

		categoryRoutes.GET("/page", categoryApi.CategoryPage)

		categoryRoutes.DELETE("/", categoryApi.CategoryDelete)

		categoryRoutes.PUT("/", categoryApi.CategoryUpdate)
	}

}
