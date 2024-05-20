package service

import (
	"github.com/gofiber/fiber/v2"
	"system_employee/model/request"
	"system_employee/model/response"
)

type CompanyEmployeeUserService interface {
	UserRegister(ctx *fiber.Ctx, params *request.UserRequest) (*fiber.Map, error)
	Login(ctx *fiber.Ctx, params *request.LoginRequest) (*response.LoginResponse, error)
}
