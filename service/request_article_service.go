package service

import (
	"github.com/gofiber/fiber/v2"
	"system_employee/model/request"
	"system_employee/model/response"
)

type RequestArticleService interface {
	InsertRequest(ctx *fiber.Ctx, params *request.ArticleRequest) (*fiber.Map, error)
	UpdateRequest(ctx *fiber.Ctx, params *request.ArticleRequest) (*fiber.Map, error)
	GetDataRequest(ctx *fiber.Ctx, params *request.GetDataRequest) (*response.GetDataArticleResponse, error)
	DeleteDataRequest(ctx *fiber.Ctx, params *request.GetDataRequest) (*fiber.Map, error)
}
