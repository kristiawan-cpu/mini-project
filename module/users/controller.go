package users

type Controller struct {
	useCase *UseCase
}

func NewController(useCase *UseCase) *Controller {
	return &Controller{
		useCase: useCase,
	}
}

type RegisterResponse struct {
	Message string           `json:"message"`
	Data    UserItemResponse `json:"data"`
}

type UserItemResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"usernme"`
	Role_ID  int    `json:"role_id" `
}

func (c Controller) Register(req *RegisterRequest) (*RegisterResponse, error) {
	user := User{Username: req.Username,
		Password: req.Password,
		RoleID:   Role{ID: 2}}
	err := c.useCase.Register(&user)
	if err != nil {
		return nil, err
	}
	registerApproval := RegisterApproval{AdminID: User{ID: user.ID}}
	if err1 := c.useCase.AddRegisterApproval(&registerApproval); err1 != nil {
		return nil, err1
	}
	res := &RegisterResponse{
		Message: "Success",
		Data: UserItemResponse{
			ID:       user.ID,
			Username: user.Username,
			Role_ID:  user.RoleID.ID,
		},
	}

	return res, nil
}

// addCustomer
type AddCustomerResponse struct {
	Masage string
	Data   CustomerItemResponse
}
type CustomerItemResponse struct {
	ID        uint   `json:"id" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Avatar    string `json:"avatar" binding:"required"`
}

func (c Controller) AddCustomer(req *AddCustomerRequest) (*AddCustomerResponse, error) {
	customer := Customer{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Avatar:    req.Avatar}
	err := c.useCase.AddCustomer(&customer)
	if err != nil {
		return nil, err
	}
	res := &AddCustomerResponse{
		Masage: "Sukses",
		Data: CustomerItemResponse{
			ID:        customer.ID,
			Firstname: customer.Firstname,
			Lastname:  customer.Lastname,
			Email:     customer.Email,
			Avatar:    customer.Avatar,
		},
	}
	return res, nil
}

type GetAllCustomerResponse struct {
	Masage string
	Data   []Customer
}

// GetAllCustomer
func (c Controller) GetAllCustomer() (*GetAllCustomerResponse, error) {
	customers, err := c.useCase.GetAllCustomer()
	if err != nil {
		return nil, err
	}
	res := &GetAllCustomerResponse{}
	for _, customer := range customers {
		item := CustomerItemResponse{
			ID:        customer.ID,
			Firstname: customer.Firstname,
			Lastname:  customer.Lastname,
			Email:     customer.Email,
			Avatar:    customer.Avatar,
		}
		res.Data = append(res.Data, Customer(item))
	}
	return res, nil

}
