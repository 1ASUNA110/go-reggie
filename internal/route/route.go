package route

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	SetupEmployeeRoutes(router)

}