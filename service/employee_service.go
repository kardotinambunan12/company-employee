package service

import (
	"github.com/gofiber/fiber/v2"
	"system_employee/model/request"
	"system_employee/model/response"
)

type CompanyEmployeeService interface {
	InsertCompany(ctx *fiber.Ctx, params *request.InsertCompanyRequest) (*fiber.Map, error)
	AdminApproval(ctx *fiber.Ctx, params *request.ApprovalRequest) (*response.ApprovalResponse, error)
	CompanyList(ctx *fiber.Ctx) (*response.ListCompanyResponse, error)
}
