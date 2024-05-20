package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	guuid "github.com/google/uuid"
	"system_employee/model/request"
	"system_employee/service"
	"system_employee/utils/logging"
)

type CompanyEmployeeUserController struct {
	CompanyEmployeeUserService service.CompanyEmployeeUserService
}

func NewCompanyEmployeeUserController(companyEmployeeUserService *service.CompanyEmployeeUserService) CompanyEmployeeUserController {
	return CompanyEmployeeUserController{CompanyEmployeeUserService: *companyEmployeeUserService}
}

func (controller *CompanyEmployeeUserController) Route(app *fiber.App) {
	app.Post("/api/user/register", controller.UserRegister)
	app.Post("/api/user/login", controller.Login)
}

func (controller *CompanyEmployeeUserController) UserRegister(ctx *fiber.Ctx) error {

	var req request.UserRequest
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

	response, err := controller.CompanyEmployeeUserService.UserRegister(ctx, &req)
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

func (controller *CompanyEmployeeUserController) Login(ctx *fiber.Ctx) error {
	var req request.LoginRequest
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

	response, err := controller.CompanyEmployeeUserService.Login(ctx, &req)
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
