package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string
	Collections []Collection
}

type Collection struct {
	gorm.Model
	UserID   uint
	Username string
	Password string
	RoleID   int
	Verified bool
	Active   bool
	User     User
}
