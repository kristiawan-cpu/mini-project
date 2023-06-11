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
	Id       uint   `json:"id"`
	Username string `json:"usernme"`
	Role_ID  uint   `json:"role_id" `
}

func (c Controller) Register(req *RegisterRequest) (*RegisterResponse, error) {
	user := User{Username: req.Username,
		Password: req.Password,
		RoleID:   2,
	}
	err := c.useCase.Register(&user)
	if err != nil {
		return nil, err
	}
	registerApproval := RegisterApproval{AdminID: user.ID}
	if err1 := c.useCase.AddRegisterApproval(&registerApproval); err1 != nil {
		return nil, err1
	}
	res := &RegisterResponse{
		Message: "Success",
		Data: UserItemResponse{
			Id:       user.ID,
			Username: user.Username,
			Role_ID:  user.RoleID},
	}

	return res, nil
}

// get all admin
type GetAllAdminResponse struct {
	GetAllAdminResponseMasage string             `json:"masage"`
	GetAllAdminResponseData   []UserItemResponse `json:"data"`
}

func (c Controller) GetAllAdmin() (*GetAllAdminResponse, error) {
	admins, err := c.useCase.GetAllAdmin()
	if err != nil {
		return nil, err
	}
	res := &GetAllAdminResponse{}
	res.GetAllAdminResponseMasage = "succes"
	for _, admin := range admins {
		item := UserItemResponse{
			Id:       admin.ID,
			Username: admin.Username,
			Role_ID:  admin.RoleID,
		}

		res.GetAllAdminResponseData = append(res.GetAllAdminResponseData, item)
	}
	return res, nil
}

// addCustomer
type AddCustomerResponse struct {
	Masage string
	Data   CustomerItemResponse
}
type CustomerItemResponse struct {
	Id        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
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
			Id:        customer.ID,
			Firstname: customer.Firstname,
			Lastname:  customer.Lastname,
			Email:     customer.Email,
			Avatar:    customer.Avatar,
		},
	}
	return res, nil
}

// GetAllCustomer
type GetAllCustomerResponse struct {
	Data []CustomerItemResponse `json:"data"`
}

func (c Controller) GetAllCustomer() (*GetAllCustomerResponse, error) {
	customers, err := c.useCase.GetAllCustomer()
	if err != nil {
		return nil, err
	}
	res := &GetAllCustomerResponse{}
	for _, customer := range customers {
		item := CustomerItemResponse{
			Id:        customer.ID,
			Firstname: customer.Firstname,
			Lastname:  customer.Lastname,
			Email:     customer.Email,
			Avatar:    customer.Avatar,
		}

		res.Data = append(res.Data, item)
	}

	return res, nil
}

// get by name and email
type GetCustomerByNameAndEmailResponse struct {
	MasageGetCustomerByNameAndEmail string               `json:"masage"`
	DataGetCustomerByNameAndEmail   CustomerItemResponse `json:"data"`
}

func (c Controller) GetCustomerByNameAndEmail(req *GetCustomerByNameAndEmailRequest) (*GetCustomerByNameAndEmailResponse, error) {
	parameterRequest := req
	customer, err := c.useCase.GetCustomerByNameAndEmail(parameterRequest)
	if err != nil {
		return nil, err
	}
	res := &GetCustomerByNameAndEmailResponse{
		MasageGetCustomerByNameAndEmail: "succes",
		DataGetCustomerByNameAndEmail: CustomerItemResponse{
			Id:        customer.ID,
			Firstname: customer.Firstname,
			Lastname:  customer.Lastname,
			Email:     customer.Email,
			Avatar:    customer.Avatar,
		},
	}
	return res, nil
}
