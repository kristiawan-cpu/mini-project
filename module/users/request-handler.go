package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestHandler struct {
	ctrl *Controller
}

func NewRequestHandler(ctrl *Controller) *RequestHandler {
	return &RequestHandler{
		ctrl: ctrl,
	}

}

func DefaultRequestHandler(db *gorm.DB) *RequestHandler {
	return NewRequestHandler(
		NewController(
			NewUseCase(
				NewRepository(db),
			),
		),
	)
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (h RequestHandler) Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	res, err := h.ctrl.Register(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

//get all admin

func (h RequestHandler) GetCustomerByNameAdnEmail(c *gin.Context) {
	res, err := h.ctrl.GetAllAdmin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// creat customer
type AddCustomerRequest struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Avatar    string `json:"avatar" binding:"required"`
}

func (h RequestHandler) AddCustomer(c *gin.Context) {
	var req AddCustomerRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	res, err := h.ctrl.AddCustomer(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// get all customer
func (h RequestHandler) GetAllCustomer(c *gin.Context) {
	res, err := h.ctrl.GetAllCustomer()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// get customer by name and email
type GetCustomerByNameAndEmailRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"firstname" binding:"required"`
}

/*
func (h RequestHandler) GetCustomerByNameAndEmail(c *gin.Context) {
	var req GetCustomerByNameAndEmailRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	res, err := h.ctrl.GetCustomerByNameAndEmail(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
*/
