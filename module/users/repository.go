package users

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (r Repository) Save(user *User) error {
	return r.db.Create(user).Error
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// find all
func (r Repository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err

}
