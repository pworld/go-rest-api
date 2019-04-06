package controller

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func TestAllEmployee(t *testing.T) {
	// Grab our router
	router := gin.Default()

	router.GET("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})

	// Convert the JSON response to a map
	var response map[string]string
	value, exists := response["data"]

	assert.Assert(t, value != "")
	assert.Assert(t, exists)
}

func TestAddEmployee(t *testing.T) {
	// Grab our router
	router := gin.Default()

	router.POST("/employee/create", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ktp":            "99912345678901",
			"first_name":     "employee",
			"last_name":      "test 1",
			"handphone":      "+6281234567890",
			"marital_status": "Single",
			"position":       "Web Developer",
			"position_desc":  "Web Developer Desc",
			"company_id":     2,
			"status_id":      1,
		})
	})

	// Convert the JSON response to a map
	var response map[string]string
	value, exists := response["data"]

	assert.Assert(t, value != "")
	assert.Assert(t, exists)
}

func TestUpdateEmployee(t *testing.T) {
	// Grab our router
	router := gin.Default()

	router.PUT("/employee/update/1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"id":             1,
			"ktp":            "99912345678901",
			"first_name":     "employee",
			"last_name":      "test 1",
			"handphone":      "+6281234567890",
			"marital_status": "Single",
			"position":       "Web Developer",
			"position_desc":  "Web Developer Desc",
			"company_id":     2,
			"status_id":      1,
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

func TestDeleteEmployee(t *testing.T) {
	// Grab our router
	router := gin.Default()

	router.DELETE("/employee/delete", func(c *gin.Context) {
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
