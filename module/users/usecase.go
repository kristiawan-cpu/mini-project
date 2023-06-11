package users

import "fmt"

type UseCase struct {
	repo *Repository
}

func NewUseCase(repo *Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

// register
func (u UseCase) Register(user *User) error {
	return u.repo.Register(user)
}

// add waiting List
func (u UseCase) AddRegisterApproval(regiterAproval *RegisterApproval) error {
	return u.repo.AddRegisterApproval(regiterAproval)
}

//alladmin
func (u UseCase) GetAllAdmin() ([]User, error) {
	return u.repo.GetAllAdmin()
}

// add customer
func (u UseCase) AddCustomer(customer *Customer) error {
	fmt.Println("case")
	return u.repo.AddCustomer(customer)
}

// get All customer
func (u UseCase) GetAllCustomer() ([]Customer, error) {
	return u.repo.GetAllCustomer()
}

/*
//Get customer by name and email
func (u UseCase) GetCustomerByNameAndEmail(getCustomerByNameAndEmail *GetCustomerByNameAndEmailRequest) (Customer, error) {
	return u.repo.GetCustomerByNameAndEmail(getCustomerByNameAndEmail)
}
*/
