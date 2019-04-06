package controller

import (
	"bytes"
	"encoding/json"
	"go-rest-api/app/model"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// AddEmployee - AddEmployee
func AddEmployee(c *gin.Context) {
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	type Emp struct {
		Ktp           string `json:"ktp" binding:"required"`
		FirstName     string `json:"first_name" binding:"required"`
		LastName      string `json:"last_name" binding:"required"`
		Handphone     string `json:"handphone" binding:"required"`
		MaritalStatus string `json:"marital_status" binding:"required"`
		Position      string `json:"position" binding:"required"`
		PositionDesc  string `json:"position_desc" binding:"required"`
		CompanyID     int    `json:"company_id" binding:"required"`
		StatusID      int    `json:"status_id" binding:"required"`
	}

	input := Emp{}
	json.Unmarshal(bodyBytes, &input)

	tx := Database.Begin()

	employee := model.Employee{
		Ktp:           input.Ktp,
		FirstName:     input.FirstName,
		LastName:      input.LastName,
		Handphone:     input.Handphone,
		MaritalStatus: input.MaritalStatus,
		LastPosition:  input.Position,
		LastCompanyId: input.CompanyID,
	}

	if err := tx.Create(&employee).Error; err != nil {
		tx.Rollback()

		c.JSON(http.StatusBadRequest, gin.H{
			"status":     http.StatusBadRequest,
			"success":    false,
			"message":    err.Error(),
			"data":       nil,
			"exceptions": nil,
		})
		return
	}

	cmpEmployee := model.CompanyEmployees{
		CompanyId:    input.CompanyID,
		EmployeeId:   employee.Id,
		StatusId:     input.StatusID,
		Position:     input.Position,
		PositionDesc: input.PositionDesc,
		StartDate:    time.Now(),
	}

	if err := tx.Create(&cmpEmployee).Error; err != nil {
		tx.Rollback()

		c.JSON(http.StatusBadRequest, gin.H{
			"status":     http.StatusBadRequest,
			"success":    false,
			"message":    err.Error(),
			"data":       nil,
			"exceptions": nil,
		})
		return
	}

	if tx.Commit().Error == nil {
		c.JSON(http.StatusCreated, gin.H{
			"status":     http.StatusCreated,
			"success":    true,
			"message":    nil,
			"data":       &cmpEmployee,
			"exceptions": nil,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":     http.StatusBadRequest,
			"success":    false,
			"message":    tx.Commit().Error,
			"data":       nil,
			"exceptions": nil,
		})
		return
	}
}

// UpdateEmployee - UpdateEmployee
func UpdateEmployee(c *gin.Context) {
	var employee model.Employee
	empID := c.Param("id")

	Database.First(&employee, empID)
	if employee.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":     http.StatusNotFound,
			"success":    false,
			"message":    "Data not found",
			"data":       nil,
			"exceptions": nil,
		})
		return
	}

	c.BindJSON(&employee)
	err := Database.Save(&employee).Error
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":     http.StatusOK,
			"success":    true,
			"message":    nil,
			"data":       &employee,
			"exceptions": nil,
		})
	} else {
		panic(err)
	}
}

// DeleteEmployee - DeleteEmployee
func DeleteEmployee(c *gin.Context) {
	var employee model.Employee
	empID := c.Param("id")

	Database.First(&employee, empID)
	if employee.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":     http.StatusNotFound,
			"success":    false,
			"message":    "Data not found",
			"data":       nil,
			"exceptions": nil,
		})
		return
	}

	err := Database.Delete(&employee).Error
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

// AllEmployee - AllEmployee
func AllEmployee(c *gin.Context) {
	var employee []model.Employee

	err := Database.Order("id asc").Find(&employee).Error

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":     http.StatusOK,
			"success":    true,
			"message":    nil,
			"data":       &employee,
			"exceptions": nil,
		})
	} else {
		panic(err)
	}
}
