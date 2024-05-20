package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	guuid "github.com/google/uuid"
	errorhandler "system_employee/error_handler"
	"system_employee/middleware"
	"system_employee/model/request"
	"system_employee/service"
	"system_employee/utils/logging"
)

type RequestemployeeController struct {
	RequestEmployeeService service.RequestArticleService
}

func NewRequestemployeeController(requestemployeeService *service.RequestArticleService) RequestemployeeController {
	return RequestemployeeController{RequestEmployeeService: *requestemployeeService}
}
func (controller *RequestemployeeController) Route(app *fiber.App) {
	app.Post("/api/article/insert", middleware.JWTProtected(), controller.InsertRequest)
	app.Post("/api/article/update", middleware.JWTProtected(), controller.UpdateRequest)
	app.Post("/api/article/get-request", middleware.JWTProtected(), controller.GetDataRequest)
	app.Post("/api/article/delete", middleware.JWTProtected(), controller.DeleteDataRequest)

}

func (controller *RequestemployeeController) InsertRequest(ctx *fiber.Ctx) error {
	var req request.ArticleRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	requestId := guuid.New()
	logStart := logging.LogRequest(ctx, string(requestData), requestId.String(), "0000")

	response, err := controller.RequestEmployeeService.InsertRequest(ctx, &req)
	if err != nil {
		return err
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		return err
	}

	logging.LogResponse(ctx, string(responseData), requestId.String(), "0000", logStart)

	return ctx.Status(200).JSON(response)
}

func (controller *RequestemployeeController) UpdateRequest(ctx *fiber.Ctx) error {
	var req request.ArticleRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return err
	}
	fmt.Println(requestData)

	requestId := guuid.New()
	logStart := logging.LogRequest(ctx, string(requestData), requestId.String(), "0000")

	response, err := controller.RequestEmployeeService.UpdateRequest(ctx, &req)
	if err != nil {
		return errors.New("gagal update data")
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		return err
	}

	logging.LogResponse(ctx, string(responseData), requestId.String(), "0000", logStart)

	return ctx.Status(200).JSON(response)
}

func (controller *RequestemployeeController) GetDataRequest(ctx *fiber.Ctx) error {
	var req request.GetDataRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	requestId := guuid.New()
	logStart := logging.LogRequest(ctx, string(requestData), requestId.String(), "0000")

	response, err := controller.RequestEmployeeService.GetDataRequest(ctx, &req)
	if err != nil {
		return err
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		return err
	}

	logging.LogResponse(ctx, string(responseData), requestId.String(), "0000", logStart)

	return ctx.Status(200).JSON(response)
}

func (controller *RequestemployeeController) DeleteDataRequest(ctx *fiber.Ctx) error {
	var req request.GetDataRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	requestId := guuid.New()
	logStart := logging.LogRequest(ctx, string(requestData), requestId.String(), "0000")

	response, err := controller.RequestEmployeeService.DeleteDataRequest(ctx, &req)
	if err != nil {
		errorhandler.PanicIfNeeded(err)
	}
	responseData, err := json.Marshal(response)
	if err != nil {
		return err
	}

	logging.LogResponse(ctx, string(responseData), requestId.String(), "0000", logStart)

	return ctx.Status(200).JSON(response)
}
