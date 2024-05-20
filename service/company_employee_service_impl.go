package service

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"system_employee/model/request"
	"system_employee/model/response"
	"system_employee/repository"
	"system_employee/utils"
	"system_employee/utils/constant"
	"system_employee/utils/logging"
)

func NewCompanyEmployeeService(companyEmployeeRepository *repository.CompanyEmployeeRepository) CompanyEmployeeService {
	return &companyEmployeeImpl{
		CompanyEmployeeUserRepository: *companyEmployeeRepository,
	}
}

type companyEmployeeImpl struct {
	CompanyEmployeeUserRepository repository.CompanyEmployeeRepository
}

func (service *companyEmployeeImpl) CompanyList(ctx *fiber.Ctx) (*response.ListCompanyResponse, error) {
	result, err := service.CompanyEmployeeUserRepository.CompanyList(ctx)
	if err != nil {
		logging.Logging(ctx, constant.ERROR_LOAD_DATA, constant.FuncGetEmployee, "0010", "", "0000", err.Error())
		return nil, errors.New("data tidak ditemukan")
	}

	resultData := &response.ListCompanyResponse{
		Message:    result.Message,
		StatusCode: result.StatusCode,
		Data:       result.Data,
	}

	return resultData, nil
}

func (service *companyEmployeeImpl) InsertCompany(ctx *fiber.Ctx, params *request.InsertCompanyRequest) (*fiber.Map, error) {
	if params.Title == "" || params.Description == "" || params.CompanyId == "" {
		message := "kolom permintaan tidak boleh kosong"
		return nil, errors.New(message)
	}

	err := service.CompanyEmployeeUserRepository.InserCompany(params)
	if err != nil {
		logging.Logging(ctx, constant.ERRORMIDDLEWARE, constant.FuncInsertEmployee, "0010", "", "0000", err.Error())
		return nil, err
	}
	result := &fiber.Map{
		"message":    "Success",
		"statusCode": 200,
	}
	return result, nil
}

func (service *companyEmployeeImpl) AdminApproval(ctx *fiber.Ctx, params *request.ApprovalRequest) (*response.ApprovalResponse, error) {
	resData, err := service.CompanyEmployeeUserRepository.GetCompanyId(ctx, params.Id)
	if err != nil {
		logging.Logging(ctx, constant.ERROR_LOAD_DATA, constant.FuncApproval, "0010", "", "0000", err.Error())
		return nil, err
	}
	if resData == "" {
		message := "Id Permintaan tidak di temukan"
		return nil, errors.New(message)
	}

	claims := utils.JwtVerificationWithClaim(ctx)
	if params.EmailUser != claims.Email {
		message := "Permintaan ditolak"
		return nil, errors.New(message)
	}

	result, err := service.CompanyEmployeeUserRepository.ApprovalStatus(params)
	if err != nil {
		logging.Logging(ctx, constant.ERROR_UPDATE_DATA, constant.FuncApproval, "0010", "", "0000", err.Error())
		return nil, err
	}

	return result, nil
}
