package users

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

// add customer
func (u UseCase) AddCustomer(customer *Customer) error {
	return u.repo.AddCustomer(customer)
}

// get All customer
func (u UseCase) GetAllCustomer() ([]Customer, error) {
	return u.repo.GetAllCustomer()
}
