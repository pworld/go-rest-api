package controller

import (
	"go-rest-api/app/model"
	"net/http"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

var identityKey = "id"

// HomePage - HomePage
func HomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":     http.StatusOK,
		"success":    true,
		"message":    nil,
		"data":       "Gin gonic framework, you know for rest api",
		"exceptions": nil,
	})
}

// HelloHandler - HelloHandler
func HelloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"success": true,
		"message": nil,
		"data": gin.H{
			"userID":   claims["id"],
			"userName": user.(*model.User).Username,
			"text":     "Hello World.",
		},
		"exceptions": nil,
	})
}
