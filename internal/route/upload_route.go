package route

import (
	"github.com/gin-gonic/gin"
	"go-reggie/internal/api"
)

func SetupUploadRoutes(router *gin.Engine) {

	uploadApi := api.NewUploadApi()

	employeeRoutes := router.Group("/upload")
	{
		employeeRoutes.POST("/", uploadApi.Upload)

	}

}
