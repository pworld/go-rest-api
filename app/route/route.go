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
		emp.PUT("/:id", controller.UpdateEmployee)
		emp.DELETE("/:id", controller.DeleteEmployee)
	}

	cmp := r.Group("/companies")
	cmp.Use(authMiddleware.MiddlewareFunc())
	{
		cmp.GET("/", controller.AllCompany)
		cmp.POST("/create", controller.CreateCompany)
		cmp.PUT("/update", controller.UpdateCompany)
		cmp.DELETE("/delete", controller.DeleteCompany)
		cmp.POST("/list", controller.ListCompany)
	}

	smp := r.Group("/employee-search")
	smp.Use(authMiddleware.MiddlewareFunc())
	{
		smp.GET("/", controller.SearchEmployee)
		smp.GET("/history", controller.SearchEmployeeHistory)
		smp.GET("/employee-company", controller.SearchEmployeeCompany)
		smp.GET("/employee-company-friends", controller.SearchEmployeeFriends)
		smp.GET("/employee-friends-friends", controller.SearchEmployeeFF)
	}

	fmp := r.Group("/employee-friend")
	fmp.Use(authMiddleware.MiddlewareFunc())
	{
		fmp.GET("/", controller.AllEmployeeFriend)
		fmp.POST("/create", controller.CreateEmployeeFriend)
		fmp.PUT("/update", controller.UpdateEmployeeFriend)
		fmp.DELETE("/delete", controller.DeleteEmployeeFriend)
	}

	cemp := r.Group("/company-employee")
	cemp.Use(authMiddleware.MiddlewareFunc())
	{
		cemp.GET("/", controller.AllCompanyEmployee)
		cemp.POST("/create", controller.CreateCompanyEmployee)
		cemp.PUT("/update", controller.UpdateCompanyEmployee)
		cemp.DELETE("/delete", controller.DeleteCompanyEmployee)
	}
	return r
}
