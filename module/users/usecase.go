package users

type UseCase struct {
	repo *Repository
}

func NewUseCase(repo *Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u UseCase) Create(user *User) error {
	return u.repo.Save(user)
}

// getall
func (u UseCase) GetAll() ([]User, error) {
	return u.repo.FindAll()
}
