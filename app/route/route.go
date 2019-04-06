package route

import (
	"go-rest-api/app/controller"
	"go-rest-api/app/middleware"

	"github.com/gin-gonic/gin"
)

// Init - Init route
func Init() *gin.Engine {
	r := gin.New()
	r.GET("/", controller.HomePage)

	authMiddleware := middleware.JwtMiddlewareHandler()

	r.POST("/login", authMiddleware.LoginHandler)

	auth := r.Group("/auth")

	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", controller.HelloHandler)
	}

	empStatus := r.Group("/employee_status")
	empStatus.Use(authMiddleware.MiddlewareFunc())
	{
		empStatus.GET("/", controller.AllEmployeeStatus)
		empStatus.GET("/:id", controller.SingleEmployeeStatus)
		empStatus.POST("/", controller.CreateEmployeeStatus)
		empStatus.PUT("/:id", controller.UpdateEmployeeStatus)
		empStatus.DELETE("/:id", controller.DeleteEmployeeStatus)
	}

	emp := r.Group("/employees")
	emp.Use(authMiddleware.MiddlewareFunc())
	{
		emp.GET("/", controller.AllEmployee)
		emp.POST("/", controller.AddEmployee)
	}

	cmp := r.Group("/companies")
	cmp.Use(authMiddleware.MiddlewareFunc())
	{
		cmp.GET("/", controller.AllCompany)
		cmp.POST("/", controller.CreateCompany)
		cmp.POST("/list", controller.ListCompany)
	}

	return r
}
