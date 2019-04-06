package model

// Employee - Employee models
type Employee struct {
	Id            int    `json:"id" orm:"auto"`
	Ktp           string `json:"ktp" orm:"size(64)"`
	FirstName     string `json:"first_name" orm:"size(64)"`
	LastName      string `json:"last_name" orm:"size(64)"`
	Handphone     string `json:"handphone" orm:"size(32)"`
	MaritalStatus string `json:"marital_status" orm:"size(64)"`
	LastPosition  string `json:"last_position" orm:"size(100)"`
	CompanyID     int    `json:"company_id" orm:"size(32)"`
}

// TableName - get table name of Employee
func (Employee) TableName() string {
	return "employees"
}

// EmployeeFriend - get data EmployeeFriend
type EmployeeFriend struct {
	ID             int `json:"id" orm:"auto"`
	EmployeeKTP    int `json:"employee_ktp" orm:"size(32)"`
	FriendKTP      int `json:"friend_ktp" orm:"size(32)"`
	EmployeeDetail Employee
	FriendDetail   Employee
}

// TableName - get table name of EmployeeFriend
func (EmployeeFriend) TableName() string {
	return "employee_friends"
}

// EmployeeStatus - EmployeeStatus models
type EmployeeStatus struct {
	Id          int    `json:"id" orm:"auto"`
	Name        string `json:"name" orm:"size(100)"`
	Description string `json:"description" orm:"size(255)"`
}

// TableName - get table name of EmployeeStatus
func (EmployeeStatus) TableName() string {
	return "employee_status"
}
