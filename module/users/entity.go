package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint
	Username string
	Password string
	RoleID   Role
	Verified bool
	Active   bool
}

type Role struct {
	ID        int
	Role_name string
}
type Customer struct {
	ID        uint
	Firstname string
	Lastname  string
	Email     string
	Avatar    string
}
type RegisterApproval struct {
	ID           User
	AdminID      User
	SuperAdminId User
	Status       string
}
