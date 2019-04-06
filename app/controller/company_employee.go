package controller

import (
	"go-rest-api/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AllCompanyEmployee - AllCompanyEmployee
func AllCompanyEmployee(c *gin.Context) {
	var companyEmployee []model.CompanyEmployees

	err := Database.Order("id asc").Find(&companyEmployee).Error

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":     http.StatusOK,
			"success":    true,
			"message":    nil,
			"data":       &companyEmployee,
			"exceptions": nil,
		})
	} else {
		panic(err)
	}
}

// CreateCompanyEmployee - CreateCompanyEmployee
func CreateCompanyEmployee(c *gin.Context) {
	var companyEmployee model.CompanyEmployees

	c.BindJSON(&companyEmployee)
	err := Database.Save(&companyEmployee).Error

	if err == nil {
		c.JSON(http.StatusCreated, gin.H{
			"status":     http.StatusCreated,
			"success":    true,
			"message":    nil,
			"data":       &companyEmployee,
			"exceptions": nil,
		})
	} else {
		panic(err)
	}
}

// UpdateCompanyEmployee - UpdateCompanyEmployee
func UpdateCompanyEmployee(c *gin.Context) {
	var companyEmployee model.CompanyEmployees
	empStatusID := c.Param("id")

	Database.First(&companyEmployee, empStatusID)
	if companyEmployee.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":     http.StatusNotFound,
			"success":    false,
			"message":    "Data not found",
			"data":       nil,
			"exceptions": nil,
		})
		return
	}

	c.BindJSON(&companyEmployee)
	err := Database.Save(&companyEmployee).Error
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":     http.StatusOK,
			"success":    true,
			"message":    nil,
			"data":       &companyEmployee,
			"exceptions": nil,
		})
	} else {
		panic(err)
	}
}

// DeleteCompanyEmployees - DeleteCompanyEmployees
func DeleteCompanyEmployee(c *gin.Context) {
	var companyEmployee model.CompanyEmployees
	empStatusID := c.Param("id")

	Database.First(&companyEmployee, empStatusID)
	if companyEmployee.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":     http.StatusNotFound,
			"success":    false,
			"message":    "Data not found",
			"data":       nil,
			"exceptions": nil,
		})
		return
	}

	err := Database.Delete(&companyEmployee).Error
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
