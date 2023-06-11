package users

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// register

func (r Repository) Register(user *User) error {
	registered := r.db.Create(&user)
	return registered.Error
}

// get all admin
func (r Repository) GetAllAdmin() ([]User, error) {
	var users []User
	if err := r.db.Find(&users); err != nil {

	}
	return users, nil
}

// addRegisterApproval
func (r Repository) AddRegisterApproval(registerApproval *RegisterApproval) error {
	return r.db.Create(registerApproval).Error
}

// addCustomer
func (r Repository) AddCustomer(cstm *Customer) error {
	return r.db.Create(cstm).Error
}

// GetAllCustomer
func (r Repository) GetAllCustomer() ([]Customer, error) {
	var customers []Customer
	if err := r.db.Find(&customers); err != nil {

	}
	return customers, nil
}

// GetCustomerByNameAdnEmail
func (r Repository) GetCustomerByNameAndEmail(req *GetCustomerByNameAndEmailRequest) (Customer, error) {

	return r.db.WithContext(ctx)
}
