package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	guuid "github.com/google/uuid"
	"system_employee/middleware"
	"system_employee/model/request"
	"system_employee/service"
	"system_employee/utils/logging"
)

type CompanyEmployeeController struct {
	CompanyEmployeeService service.CompanyEmployeeService
}

func NewCompanyController(companyEmployeeService *service.CompanyEmployeeService) CompanyEmployeeController {
	return CompanyEmployeeController{CompanyEmployeeService: *companyEmployeeService}
}

func (controller *CompanyEmployeeController) Route(app *fiber.App) {
	app.Post("/api/company/create", controller.InsertCompany)
	app.Get("/api/company/list", controller.CompanyList)
	app.Post("/api/admin/approval", middleware.JWTProtected(), controller.AdminApproval)

}

func (controller *CompanyEmployeeController) InsertCompany(ctx *fiber.Ctx) error {
	var req request.InsertCompanyRequest
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

	response, err := controller.CompanyEmployeeService.InsertCompany(ctx, &req)
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

func (controller *CompanyEmployeeController) AdminApproval(ctx *fiber.Ctx) error {
	var req request.ApprovalRequest
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

	response, err := controller.CompanyEmployeeService.AdminApproval(ctx, &req)
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

func (controller *CompanyEmployeeController) CompanyList(ctx *fiber.Ctx) error {
	requestId := guuid.New()
	logStart := logging.LogRequest(ctx, string("Employee List"), requestId.String(), "0000")

	response, err := controller.CompanyEmployeeService.CompanyList(ctx)
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
