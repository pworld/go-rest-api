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

	return r
}
