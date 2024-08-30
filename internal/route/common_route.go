package route

import (
	"github.com/gin-gonic/gin"
	"go-reggie/internal/api"
)

func SetupCommonRoutes(router *gin.Engine) {

	uploadApi := api.NewCommonApi()

	employeeRoutes := router.Group("/common")
	{
		employeeRoutes.POST("/upload", uploadApi.FileUpload)

		employeeRoutes.GET("/download", uploadApi.FileDownload)

	}

}
