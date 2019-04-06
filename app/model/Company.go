package model

import "time"

// Company - Company models
type Company struct {
	Id               int    `json:"id" orm:"auto"`
	Name             string `json:"name" orm:"size(100)"`
	TDP              string `json:"tdp" orm:"size(64)"`
	Email            string `json:"email" orm:"size(64)"`
	PrimaryContact   string `json:"primary_contact" orm:"size(100)"`
	Phone            string `json:"phone" orm:"size(32)"`
	Address          string `json:"address" orm:"size(100)"`
	EmpCount         int    `json:"emp_count" orm:"size(32)"`
	CompanyEmployees []CompanyEmployees
}

// TableName - get table name of Company
func (Company) TableName() string {
	return "companies"
}

// CompanyEmployees - CompanyEmployees models
type CompanyEmployees struct {
	Id           int    `json:"id" orm:"auto"`
	CompanyId    int    `json:"company_id" orm:"size(32)"`
	EmployeeId   int    `json:"employee_id" orm:"size(32)"`
	StatusId     int    `json:"status_id" orm:"size(32)"`
	Position     string `json:"position" orm:"size(100)"`
	PositionDesc string `json:"position_desc" orm:"size(255)"`
	StartDate    time.Time
	EndDate      time.Time
}

// TableName - get table name of company_employees
func (CompanyEmployees) TableName() string {
	return "company_employees"
}
