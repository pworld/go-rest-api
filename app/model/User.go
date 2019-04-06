package model

import "time"

// User - user models
type User struct {
	ID          int    `json:"id" orm:"auto"`
	Username    string `json:"username" orm:"size(64)"`
	Email       string `json:"email" orm:"size(64)"`
	Password    string `json:"password" orm:"size(100)"`
	VerifyAt    time.Time
	LastLoginAt time.Time
}

// TableName - tabelname
func (User) TableName() string {
	return "users"
}
