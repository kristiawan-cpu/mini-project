package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	RoleID   uint
	//	Verified bool
	//	Active   bool
}

type Role struct {
	Id        uint `gorm:"primary_key"`
	Role_name string
}
type Customer struct {
	gorm.Model
	Firstname string
	Lastname  string
	Email     string
	Avatar    string
}
type RegisterApproval struct {
	Id           uint `gorm:"primary_key"`
	AdminID      uint
	SuperAdminId uint
	Status       string
}
