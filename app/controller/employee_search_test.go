package controller

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func TestSearchEmployeeHistory(t *testing.T) {
	// Grab our router
	router := gin.Default()

	router.GET("/history", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"tdp":       "123",
			"status_id": 1,
		})
	})

	// Convert the JSON response to a map
	var response map[string]string
	value, exists := response["data"]

	assert.Assert(t, value != "")
	assert.Assert(t, exists)
}

func TestSearchEmployee(t *testing.T) {
	// Grab our router
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "123",
			"ktp":  1,
		})
	})

	// Convert the JSON response to a map
	var response map[string]string
	value, exists := response["data"]

	assert.Assert(t, value != "")
	assert.Assert(t, exists)
}

func TestSearchEmployeeCompany(t *testing.T) {
	// Grab our router
	router := gin.Default()

	router.GET("/employee-company", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ktp": 1,
		})
	})

	// Convert the JSON response to a map
	var response map[string]string
	value, exists := response["data"]

	assert.Assert(t, value != "")
	assert.Assert(t, exists)
}

func TestSearchEmployeeFriends(t *testing.T) {
	// Grab our router
	router := gin.Default()

	router.GET("/employee-company-friends", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ktp": 1,
		})
	})

	// Convert the JSON response to a map
	var response map[string]string
	value, exists := response["data"]

	assert.Assert(t, value != "")
	assert.Assert(t, exists)
}

func TestSearchEmployeeCompanyFriends(t *testing.T) {
	// Grab our router
	router := gin.Default()

	router.GET("/employee-friends-friends", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ktp": 1,
		})
	})

	// Convert the JSON response to a map
	var response map[string]string
	value, exists := response["data"]

	assert.Assert(t, value != "")
	assert.Assert(t, exists)
}
