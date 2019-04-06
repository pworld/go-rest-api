package controller

import (
	"bytes"
	"encoding/json"
	"go-rest-api/app/model"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SearchEmployeeHistory - SearchEmployeeHistory
/*
 * c.Mendapatkan daftar karyawan yang SEDANG bekerja di perusahaan berdasarkan TDP
 * d.Mendapatkan daftar karyawan yang TELAH/SEDANG bekerja di perusahaan berdasarkan TDP (request status_id = 0)
 */
func SearchEmployeeHistory(c *gin.Context) {
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	type Emp struct {
		TDP      string `json:"tdp" binding:"required"`
		StatusID int    `json:"status_id" binding:"required"`
	}

	input := Emp{}
	json.Unmarshal(bodyBytes, &input)

	var employee []model.Employee
	query := Database.Table("companies")
	query = query.Joins("left join company_employees on company_employees.company_id = companies.id")
	query = query.Joins("left join employees on employees.id = company_employees.employee_id")

	if input.TDP != "" {
		query = query.Where("companies.tdp LIKE ?", "%"+input.TDP+"%")
	}
	if input.StatusID != 0 {
		switch input.StatusID {
		case 1:
			query = query.Where("company_employees.status_id = 1")
		case 2:
			query = query.Where("company_employees.status_id = 2")
		}
	}

	err := query.Find(&employee).Error

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":     http.StatusOK,
			"success":    true,
			"message":    nil,
			"data":       err,
			"exceptions": nil,
		})
	} else {
		panic(err)
	}
}

/*
 * SearchEmployee- SearchEmployee
 * e.Mencari karyawan berdasarkan nama
 * f.Mendapatkan informasi karyawan berdasarkan KTP karyawan
 */
func SearchEmployee(c *gin.Context) {
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	type Filter struct {
		Name string `json:"name" binding:"required"`
		Ktp  string `json:"ktp" binding:"required"`
	}

	input := Filter{}
	json.Unmarshal(bodyBytes, &input)

	var employee []model.Employee
	query := Database.Table("employees")

	if input.Name != "" {
		query = query.Where("employees.name LIKE ?", "%"+input.Name+"%")
	}
	if input.Ktp != "" {
		query = query.Where("employees.ktp LIKE ?", "%"+input.Ktp+"%")
	}

	err := query.Find(&employee).Error

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":     http.StatusOK,
			"success":    true,
			"message":    nil,
			"data":       err,
			"exceptions": nil,
		})
	} else {
		panic(err)
	}
}

// SearchEmployeeCompany- SearchEmployeeCompany
/*
 * g. Mendapatkan riwayat pekerjaan karyawan berdasarkan KTP karyawan.
 */
func SearchEmployeeCompany(c *gin.Context) {
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	type Filter struct {
		Ktp string `json:"ktp" binding:"required"`
	}

	input := Filter{}
	json.Unmarshal(bodyBytes, &input)

	var employee []model.Employee
	query := Database.Table("employees")
	query = query.Joins("left join company_employees on company_employees.employee_id = employees.id")
	query = query.Joins("left join companies on companies.id = company_employees.company_id")

	if input.Ktp != "" {
		query = query.Where("employees.ktp = ?", input.Ktp)
	}

	err := query.Find(&employee).Error

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":     http.StatusOK,
			"success":    true,
			"message":    nil,
			"data":       err,
			"exceptions": nil,
		})
	} else {
		panic(err)
	}
}

// SearchEmployeeFriends - SearchEmployeeFriends
/*
 * h. Mendapatkan karyawan-karyawan lain yang menjadi teman dari seorang karyawan (berdasarkan KTP)
 */
func SearchEmployeeFriends(c *gin.Context) {
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	type Filter struct {
		Ktp string `json:"ktp" binding:"required"`
	}

	input := Filter{}
	json.Unmarshal(bodyBytes, &input)

	var employee []model.Employee
	if input.Ktp != "" {
		// * sub query to handle process
		query := Database.Where("employees.ktp = ?", Database.Table("employee_friends").Select("employee_friends.friend_ktp").Where("employee_friends.employee_ktp = ?", input.Ktp).QueryExpr())
		err := query.Find(&employee).Error
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status":     http.StatusOK,
				"success":    true,
				"message":    nil,
				"data":       err,
				"exceptions": nil,
			})
		} else {
			panic(err)
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":     http.StatusOK,
			"success":    false,
			"message":    nil,
			"data":       "Empty Data",
			"exceptions": nil,
		})
	}
}

// SearchEmployeeCompanyFriends - SearchEmployeeCompanyFriends
/*
 * i. Mendapatkan karyawan-karyawan lain yang dianggap pernah bekerja bersama-sama** dengan seorang karyawan (berdasarkan KTP)
 */
func SearchEmployeeCompanyFriends(c *gin.Context) {
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	type Filter struct {
		Ktp string `json:"ktp" binding:"required"`
	}

	input := Filter{}
	json.Unmarshal(bodyBytes, &input)

	var employee []model.Employee
	if input.Ktp != "" {
		// * sub query to handle process
		subQuery := Database.Table("employees").Select("company_employees.friend_ktp")
		subQuery.Joins("left join company_employees on company_employees.employee_id = employees.id")
		subQuery.Where("employees.employee_ktp = ?", input.Ktp).QueryExpr()

		query := Database.Where("employees.employee_ktp = ?", subQuery)
		err := query.Find(&employee).Error
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status":     http.StatusOK,
				"success":    true,
				"message":    nil,
				"data":       err,
				"exceptions": nil,
			})
		} else {
			panic(err)
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":     http.StatusOK,
			"success":    false,
			"message":    nil,
			"data":       "Empty Data",
			"exceptions": nil,
		})
	}
}

// SearchEmployeeFF - SearchEmployeeFF
/*
 * j. Mencari teman-temannya teman-teman, yang bukan teman-teman** dari seorang-karyawan (berdasarkan KTP).
 */
func SearchEmployeeFF(c *gin.Context) {
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	type Filter struct {
		Ktp string `json:"ktp" binding:"required"`
	}

	input := Filter{}
	json.Unmarshal(bodyBytes, &input)

	var employee []model.Employee
	if input.Ktp != "" {
		// * sub query to handle process
		subQuery := Database.Table("employees").Select("company_employees.friend_ktp")
		subQuery.Where("employees.employee_ktp = ?", input.Ktp).QueryExpr()

		query := Database.Where("employees.amount != ?", subQuery)
		err := query.Find(&employee).Error
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status":     http.StatusOK,
				"success":    true,
				"message":    nil,
				"data":       err,
				"exceptions": nil,
			})
		} else {
			panic(err)
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":     http.StatusOK,
			"success":    false,
			"message":    nil,
			"data":       "Empty Data",
			"exceptions": nil,
		})
	}
}
