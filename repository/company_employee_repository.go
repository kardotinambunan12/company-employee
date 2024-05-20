package repository

import (
	"github.com/gofiber/fiber/v2"
	"system_employee/model/request"
	"system_employee/model/response"
)

type CompanyEmployeeRepository interface {
	InserCompany(params *request.InsertCompanyRequest) error
	ApprovalStatus(params *request.ApprovalRequest) (*response.ApprovalResponse, error)
	CompanyList(ctx *fiber.Ctx) (*response.ListCompanyResponse, error)
	GetCompanyId(ctx *fiber.Ctx, id int) (string, error)
}
