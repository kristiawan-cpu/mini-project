package users

type Controller struct {
	useCase *UseCase
}

func NewController(useCase *UseCase) *Controller {
	return &Controller{
		useCase: useCase,
	}
}

type CreateResponse struct {
	Message string           `json:"message"`
	Data    UserItemResponse `json:"data"`
}

type CollectionItemResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UserItemResponse struct {
	ID          uint                     `json:"id"`
	Name        string                   `json:"name"`
	Collections []CollectionItemResponse `json:"collections"`
}

func (c Controller) Create(req *CreateRequest) (*CreateResponse, error) {
	user := User{Name: req.Name}
	err := c.useCase.Create(&user)
	if err != nil {
		return nil, err
	}

	res := &CreateResponse{
		Message: "Success",
		Data: UserItemResponse{
			ID:   user.ID,
			Name: user.Name,
		},
	}

	return res, nil
}

type GetResponse struct {
	Data []UserItemResponse `json:"data"`
}

func (c Controller) GetAll() (*GetResponse, error) {
	users, err := c.useCase.GetAll()
	if err != nil {
		return nil, err
	}

	res := &GetResponse{}
	for _, user := range users {
		item := UserItemResponse{
			ID:   user.ID,
			Name: user.Name,
		}
		for _, collection := range user.Collections {
			item.Collections = append(item.Collections, CollectionItemResponse{
				ID:   collection.ID,
				Name: collection.Username,
			})
		}
		res.Data = append(res.Data, item)
	}

	return res, nil
}
