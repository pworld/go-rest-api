package controller

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func TestAllEmployeeStatus(t *testing.T) {
	// Grab our router
	router := gin.Default()

	router.GET("/employee_status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})

	// Convert the JSON response to a map
	var response map[string]string
	value, exists := response["data"]

	assert.Assert(t, value != "")
	assert.Assert(t, exists)
}

func TestCreateEmployeeStatus(t *testing.T) {
	// Grab our router
	router := gin.Default()

	router.POST("/employee_status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":            "PT SENTOSA",
			"tdp":             "STS12345",
			"email":           "admin@sentosa.com",
			"primary_contact": "Bpk Sentosa",
			"phone":           "+6281268900",
			"address":         "test address",
			"emp_count":       10,
		})
	})

	// Convert the JSON response to a map
	var response map[string]string
	value, exists := response["data"]

	assert.Assert(t, value != "")
	assert.Assert(t, exists)
}

func TestUpdateEmployeeStatus(t *testing.T) {
	// Grab our router
	router := gin.Default()

	router.POST("/employee_status/1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":        "Resigned 2",
			"description": "Resigned 2",
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

func TestDeleteEmployeeStatus(t *testing.T) {
	// Grab our router
	router := gin.Default()

	router.POST("/employee_status/delete", func(c *gin.Context) {
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
