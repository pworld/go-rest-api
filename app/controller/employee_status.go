package controller

import (
	"go-rest-api/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AllEmployeeStatus - AllEmployeeStatus
func AllEmployeeStatus(c *gin.Context) {
	var employeeStatus []model.EmployeeStatus

	err := Database.Order("id asc").Find(&employeeStatus).Error

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":     http.StatusOK,
			"success":    true,
			"message":    nil,
			"data":       &employeeStatus,
			"exceptions": nil,
		})
	} else {
		panic(err)
	}
}

// SingleEmployeeStatus - SingleEmployeeStatus
func SingleEmployeeStatus(c *gin.Context) {
	var employeeStatus model.EmployeeStatus
	empStatusID := c.Param("id")

	Database.First(&employeeStatus, empStatusID)
	if employeeStatus.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":     http.StatusNotFound,
			"success":    false,
			"message":    "Data not found",
			"data":       nil,
			"exceptions": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     http.StatusOK,
		"success":    true,
		"message":    nil,
		"data":       &employeeStatus,
		"exceptions": nil,
	})
}

// CreateEmployeeStatus - CreateEmployeeStatus
func CreateEmployeeStatus(c *gin.Context) {
	var employeeStatus model.EmployeeStatus

	c.BindJSON(&employeeStatus)
	err := Database.Save(&employeeStatus).Error

	if err == nil {
		c.JSON(http.StatusCreated, gin.H{
			"status":     http.StatusCreated,
			"success":    true,
			"message":    nil,
			"data":       &employeeStatus,
			"exceptions": nil,
		})
	} else {
		panic(err)
	}
}

// UpdateEmployeeStatus - UpdateEmployeeStatus
func UpdateEmployeeStatus(c *gin.Context) {
	var employeeStatus model.EmployeeStatus
	empStatusID := c.Param("id")

	Database.First(&employeeStatus, empStatusID)
	if employeeStatus.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":     http.StatusNotFound,
			"success":    false,
			"message":    "Data not found",
			"data":       nil,
			"exceptions": nil,
		})
		return
	}

	c.BindJSON(&employeeStatus)
	err := Database.Save(&employeeStatus).Error
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":     http.StatusOK,
			"success":    true,
			"message":    nil,
			"data":       &employeeStatus,
			"exceptions": nil,
		})
	} else {
		panic(err)
	}
}

// DeleteEmployeeStatus - DeleteEmployeeStatus
func DeleteEmployeeStatus(c *gin.Context) {
	var employeeStatus model.EmployeeStatus
	empStatusID := c.Param("id")

	Database.First(&employeeStatus, empStatusID)
	if employeeStatus.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":     http.StatusNotFound,
			"success":    false,
			"message":    "Data not found",
			"data":       nil,
			"exceptions": nil,
		})
		return
	}

	err := Database.Delete(&employeeStatus).Error
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":     http.StatusOK,
			"success":    true,
			"message":    "Employee status deleted successfully",
			"data":       nil,
			"exceptions": nil,
		})
	} else {
		panic(err)
	}
}
