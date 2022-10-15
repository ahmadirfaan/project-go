package web

type RegisterCustomerRequest struct {
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required,min=8,max=100"`
	Name        string `json:"name" validate:"required"`
	NoHandphone string `json:"noHandphone" validate:"required,min=9,max=12,numeric"`
}
