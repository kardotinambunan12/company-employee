package service

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"system_employee/model/request"
	"system_employee/model/response"
	"system_employee/repository"
	"system_employee/utils"
)

func NewCompanyEmployeeUserService(companyEmployeeUserRepository *repository.CompanyEmployeeUserRepository) CompanyEmployeeUserService {
	return &companyEmployeeUserImpl{
		CompanyEmployeeUserRepository: *companyEmployeeUserRepository,
	}
}

type companyEmployeeUserImpl struct {
	CompanyEmployeeUserRepository repository.CompanyEmployeeUserRepository
}

func (service *companyEmployeeUserImpl) UserRegister(ctx *fiber.Ctx, params *request.UserRequest) (*fiber.Map, error) {
	var result *fiber.Map

	getEmailRequest := &request.GetEmailRequest{
		Email: params.Email,
	}
	//check email
	checkEmail := service.CompanyEmployeeUserRepository.GetEmail(getEmailRequest)
	fmt.Println("len ", len(checkEmail.Email))

	//check apakah email sudah pernah digunakan
	if checkEmail.Email != "" {
		message := "Email sudah pernah digunakan, silahkan gunakan email lain"
		return nil, errors.New(message)
	}

	dataType := request.DataType{
		RoleUser:  params.Data.RoleUser,
		CompanyId: params.Data.CompanyId,
	}
	aktivasiAkunRequest := &request.UserRequest{
		Name:     params.Name,
		Address:  params.Address,
		Email:    params.Email,
		Password: utils.ConvertToSHA1(params.Password),
		Data:     dataType,
	}

	err1 := service.CompanyEmployeeUserRepository.UserRegister(aktivasiAkunRequest)
	if err1 != nil {
		return nil, err1
	}

	result = &fiber.Map{
		"isSuccessful": true,
		"message":      "Success",
	}

	return result, nil
}

func (service *companyEmployeeUserImpl) Login(ctx *fiber.Ctx, params *request.LoginRequest) (*response.LoginResponse, error) {
	params.Password = utils.ConvertToSHA1(params.Password)

	if params.Password == "" && params.Email == "" {
		message := "kolom permintaan tidak boleh kosong"
		return nil, errors.New(message)
	}

	result, err := service.CompanyEmployeeUserRepository.Login(params)
	if err != nil {
		return nil, err
	}

	accessToken, err := utils.GenerateNewAccessToken(result.Email)
	if err != nil {
		return nil, err
	}

	if result.Message == "success" {
		resultLogin := &response.LoginResponse{
			Message:      "Success",
			IsSuccessful: true,
			AccessToken:  accessToken,
		}
		return resultLogin, err
	} else {
		return nil, errors.New("user login tidak di temukan")
	}
}
