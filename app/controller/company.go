package controller

import (
	"bytes"
	"encoding/json"
	"go-rest-api/app/model"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AllCompany - AllCompany
func AllCompany(c *gin.Context) {
	var company []model.Company

	err := Database.Order("id asc").Find(&company).Error

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":     http.StatusOK,
			"success":    true,
			"message":    nil,
			"data":       &company,
			"exceptions": nil,
		})
	} else {
		panic(err)
	}
}

// CreateCompany - CreateCompany
func CreateCompany(c *gin.Context) {
	var company model.Company

	c.BindJSON(&company)
	err := Database.Save(&company).Error

	if err == nil {
		c.JSON(http.StatusCreated, gin.H{
			"status":     http.StatusCreated,
			"success":    true,
			"message":    nil,
			"data":       &company,
			"exceptions": nil,
		})
	} else {
		panic(err)
	}
}

// UpdateCompany - UpdateCompany
func UpdateCompany(c *gin.Context) {
	var company model.Company
	empID := c.Param("id")

	Database.First(&company, empID)
	if company.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":     http.StatusNotFound,
			"success":    false,
			"message":    "Data not found",
			"data":       nil,
			"exceptions": nil,
		})
		return
	}

	c.BindJSON(&company)
	err := Database.Save(&company).Error
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":     http.StatusOK,
			"success":    true,
			"message":    nil,
			"data":       &company,
			"exceptions": nil,
		})
	} else {
		panic(err)
	}
}

// DeleteCompany - DeleteCompany
func DeleteCompany(c *gin.Context) {
	var company model.Company
	empID := c.Param("id")

	Database.First(&company, empID)
	if company.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":     http.StatusNotFound,
			"success":    false,
			"message":    "Data not found",
			"data":       nil,
			"exceptions": nil,
		})
		return
	}

	err := Database.Delete(&company).Error
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":     http.StatusOK,
			"success":    true,
			"message":    "Employee deleted successfully",
			"data":       nil,
			"exceptions": nil,
		})
	} else {
		panic(err)
	}
}

// ListCompany - ListCompany
/*
 * a. Mencari perusahaan berdasarkan nama perusahaan
 * b. Mendapatkan informasi perusahaan berdasarkan TDP
 */
func ListCompany(c *gin.Context) {
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	type Filter struct {
		Name string `json:"name" binding:"required"`
		Tdp  string `json:"tdp" binding:"required"`
	}

	input := Filter{}
	json.Unmarshal(bodyBytes, &input)

	var companies []model.Company
	query := Database.Table("companies")
	if input.Name != "" {
		query = query.Where("name LIKE ?", "%"+input.Name+"%")
	}
	if input.Tdp != "" {
		query = query.Where("tdp LIKE ?", "%"+input.Tdp+"%")
	}

	err := query.Find(&companies).Error
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":     http.StatusOK,
			"success":    true,
			"message":    nil,
			"data":       &companies,
			"exceptions": nil,
		})
	} else {
		panic(err)
	}
}
