package controller

import (
	"go-rest-api/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AllEmployeeFriend - AllEmployeeFriend
func AllEmployeeFriend(c *gin.Context) {
	var employeeFriend []model.EmployeeFriend

	err := Database.Order("id asc").Find(&employeeFriend).Error

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":     http.StatusOK,
			"success":    true,
			"message":    nil,
			"data":       &employeeFriend,
			"exceptions": nil,
		})
	} else {
		panic(err)
	}
}

// CreateEmployeeFriend - CreateEmployeeFriend
func CreateEmployeeFriend(c *gin.Context) {
	var employeeFriend model.EmployeeFriend

	c.BindJSON(&employeeFriend)
	err := Database.Save(&employeeFriend).Error

	if err == nil {
		c.JSON(http.StatusCreated, gin.H{
			"status":     http.StatusCreated,
			"success":    true,
			"message":    nil,
			"data":       &employeeFriend,
			"exceptions": nil,
		})
	} else {
		panic(err)
	}
}

// UpdateEmployeeFriend - UpdateEmployeeFriend
func UpdateEmployeeFriend(c *gin.Context) {
	var employeeFriend model.EmployeeFriend
	empStatusID := c.Param("id")

	Database.First(&employeeFriend, empStatusID)
	if employeeFriend.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":     http.StatusNotFound,
			"success":    false,
			"message":    "Data not found",
			"data":       nil,
			"exceptions": nil,
		})
		return
	}

	c.BindJSON(&employeeFriend)
	err := Database.Save(&employeeFriend).Error
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":     http.StatusOK,
			"success":    true,
			"message":    nil,
			"data":       &employeeFriend,
			"exceptions": nil,
		})
	} else {
		panic(err)
	}
}

// DeleteEmployeeFriend - DeleteEmployeeFriend
func DeleteEmployeeFriend(c *gin.Context) {
	var employeeFriend model.EmployeeFriend
	empStatusID := c.Param("id")

	Database.First(&employeeFriend, empStatusID)
	if employeeFriend.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":     http.StatusNotFound,
			"success":    false,
			"message":    "Data not found",
			"data":       nil,
			"exceptions": nil,
		})
		return
	}

	err := Database.Delete(&employeeFriend).Error
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
