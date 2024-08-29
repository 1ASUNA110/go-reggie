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

		employeeRoutes.GET("/page", employeeApi.EmployeePage)

		employeeRoutes.POST("/", employeeApi.EmployeeSave)

		employeeRoutes.PUT("/", employeeApi.EmployeeUpdate)

		employeeRoutes.GET("/:id", employeeApi.EmployeeGetById)

	}

}
