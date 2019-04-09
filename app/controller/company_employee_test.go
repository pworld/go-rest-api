package controller

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func TestAllCompanyEmployees(t *testing.T) {
	// Grab our router
	router := gin.Default()

	router.GET("/company-employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})

	// Convert the JSON response to a map
	var response map[string]string
	value, exists := response["data"]

	assert.Assert(t, value != "")
	assert.Assert(t, exists)
}

func TestCreateCompanyEmployees(t *testing.T) {
	// Grab our router
	router := gin.Default()

	router.POST("/company-employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"company_id":    1,
			"employee_id":   1,
			"status_id":     1,
			"position":      "123",
			"position_desc": "123",
		})
	})

	// Convert the JSON response to a map
	var response map[string]string
	value, exists := response["data"]

	assert.Assert(t, value != "")
	assert.Assert(t, exists)
}

func TestUpdateCompanyEmployees(t *testing.T) {
	// Grab our router
	router := gin.Default()

	router.POST("/company-employee/1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"id":            1,
			"company_id":    1,
			"employee_id":   1,
			"status_id":     1,
			"position":      "123",
			"position_desc": "123",
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

func TestDeleteCompanyEmployees(t *testing.T) {
	// Grab our router
	router := gin.Default()

	router.POST("/company-employee/delete", func(c *gin.Context) {
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
