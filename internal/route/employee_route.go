package route

import (
	"github.com/gin-gonic/gin"
	"go-reggie/internal/api"
)

func SetupEmployeeRoutes(router *gin.Engine) {

	employeeApi := api.NewEmployeeApi()

	employeeRoutes := router.Group("/employee")
	{
		employeeRoutes.POST("/login", employeeApi.EmployeeLogin)

		employeeRoutes.POST("/logout", employeeApi.EmployeeLogout)

		employeeRoutes.POST("/", employeeApi.EmployeeSave)
		// 其他 employee 路由
	}

}
