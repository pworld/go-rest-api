package controller

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func TestAllEmployeeFriend(t *testing.T) {
	// Grab our router
	router := gin.Default()

	router.GET("/employee-friend", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})

	// Convert the JSON response to a map
	var response map[string]string
	value, exists := response["data"]

	assert.Assert(t, value != "")
	assert.Assert(t, exists)
}

func TestCreateEmployeeFriend(t *testing.T) {
	// Grab our router
	router := gin.Default()

	router.POST("/employee-friend", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"employee_ktp": "123",
			"friend_ktp":   "123",
		})
	})

	// Convert the JSON response to a map
	var response map[string]string
	value, exists := response["data"]

	assert.Assert(t, value != "")
	assert.Assert(t, exists)
}

func TestUpdateEmployeeFriend(t *testing.T) {
	// Grab our router
	router := gin.Default()

	router.POST("/employee-friend/1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"employee_ktp": "123",
			"friend_ktp":   "123",
		})
	})

	// Convert the JSON response to a map
	var response map[string]string
	value, exists := response["data"]

	// Failed Assert
	// assert.Assert(t, value["name"] != "PT ABC")
	assert.Assert(t, value != "")
	assert.Assert(t, exists)
}

func TestDeleteEmployeeFriend(t *testing.T) {
	// Grab our router
	router := gin.Default()

	router.POST("/employee-friend/delete", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"id": 1,
		})
	})

	// Convert the JSON response to a map
	var response map[string]string
	value, exists := response["data"]

	// Failed Assert
	assert.Assert(t, value == "")
	assert.Assert(t, exists)
}
