package repository

import (
	"github.com/gofiber/fiber/v2"
	"system_employee/model"
	"system_employee/model/request"
	"system_employee/model/response"
)

type RequestArticleRepository interface {
	InsertRequest(params *request.ArticleRequest) error
	UpdateRequest(params *request.ArticleRequest) error
	GetDataRequest(ctx *fiber.Ctx, email string, restDataAdmin string) (*response.GetDataArticleResponse, error)
	DeleteDataRequest(id int) (*model.GeneralResponse, error)
	CheckIdRequest(id int) (int, error)
	GetCompanyId(ctx *fiber.Ctx, email string) (string, error)
}
