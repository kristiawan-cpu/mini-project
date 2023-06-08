package users

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

//register
func (r Repository) Register(user *User) error {
	return r.db.Create(user).Error
}

//addRegisterApproval
func (r Repository) AddRegisterApproval(registerApproval *RegisterApproval) error {
	return r.db.Create(registerApproval).Error
}

//addCustomer
func (r Repository) AddCustomer(customer *Customer) error {
	return r.db.Create(customer).Error
}

//GetAllCustomer
func (r Repository) GetAllCustomer() ([]Customer, error) {
	var customers []Customer
	err := r.db.Find(&customers).Error
	return customers, err
}
