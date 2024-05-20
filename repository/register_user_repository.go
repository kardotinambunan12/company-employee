package repository

import (
	"system_employee/model/request"
	"system_employee/model/response"
)

type CompanyEmployeeUserRepository interface {
	GetEmail(request *request.GetEmailRequest) *response.EmailResponse
	UserRegister(params *request.UserRequest) error
	Login(params *request.LoginRequest) (*response.LoginResponse, error)
}
