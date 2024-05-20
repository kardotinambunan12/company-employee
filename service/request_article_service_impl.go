package service

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"system_employee/model/request"
	"system_employee/model/response"
	"system_employee/repository"
	"system_employee/utils"
	"system_employee/utils/constant"
	"system_employee/utils/logging"
)

func NewRequestArticleService(requestEmployeeRepository *repository.RequestArticleRepository) RequestArticleService {
	return &requestArticleImpl{
		RequestEmployeeRepository: *requestEmployeeRepository,
	}
}

type requestArticleImpl struct {
	RequestEmployeeRepository repository.RequestArticleRepository
}

func (service *requestArticleImpl) InsertRequest(ctx *fiber.Ctx, params *request.ArticleRequest) (*fiber.Map, error) {
	claims := utils.JwtVerificationWithClaim(ctx)
	if params.Title == "" || params.EmailUser == "" || params.Description == "" || params.Data.EmployeeId == "" {
		message := "kolom permintaan tidak boleh kosong"
		return nil, errors.New(message)
	}

	if params.EmailUser != claims.Email {
		message := "Permintaan ditolak"
		return nil, errors.New(message)
	}

	if claims.Role != "Karyawan" {
		message := "Akses kamu untuk fitur ini ditolak"
		return nil, errors.New(message)
	}

	err := service.RequestEmployeeRepository.InsertRequest(params)
	if err != nil {
		logging.Logging(ctx, constant.ERROR_INSERT_DATA, constant.FuncInsertEmployee, "0010", "", "0000", err.Error())
		return nil, errors.New("gagal melakukan insert data")
	}
	result := &fiber.Map{
		"message":    "Success",
		"statusCode": 200,
	}
	return result, nil
}

func (service *requestArticleImpl) UpdateRequest(ctx *fiber.Ctx, params *request.ArticleRequest) (*fiber.Map, error) {
	if params.EmailUser == "" {
		message := "Email tidak boleh kosong"
		return nil, errors.New(message)
	}

	claims := utils.JwtVerificationWithClaim(ctx)
	if params.EmailUser != claims.Email {
		message := "Permintaan ditolak"
		return nil, errors.New(message)
	}
	if claims.Role != "Karyawan" {
		message := "Akses kamu untuk fitur ini ditolak"
		return nil, errors.New(message)
	}
	fmt.Println("params : ", params)

	fmt.Println("claims : ", claims)

	err := service.RequestEmployeeRepository.UpdateRequest(params)
	if err != nil {
		logging.Logging(ctx, constant.ERROR_UPDATE_DATA, constant.FuncUpdateEmployee, "0010", "", "0000", err.Error())
		return nil, errors.New("gagal melakukan update data")
	}
	result := &fiber.Map{
		"message":    "Success",
		"statusCode": 200,
	}
	return result, nil
}

func (service *requestArticleImpl) GetDataRequest(ctx *fiber.Ctx, params *request.GetDataRequest) (*response.GetDataArticleResponse, error) {
	var (
		requestData  *response.GetDataArticleResponse
		resDataAdmin string
	)
	if params.EmailUser == "" {
		message := "Email tidak boleh kosong"
		return nil, errors.New(message)
	}

	claims := utils.JwtVerificationWithClaim(ctx)
	if params.EmailUser != claims.Email {
		message := "Permintaan ditolak"
		return nil, errors.New(message)
	}
	if claims.Role != "Karyawan" && claims.Role != "Admin" {
		message := "Akses kamu untuk fitur ini ditolak"
		return nil, errors.New(message)
	}

	if claims.Role == "Karyawan" {
		result, err := service.RequestEmployeeRepository.GetDataRequest(ctx, params.EmailUser, resDataAdmin)
		if err != nil {
			logging.Logging(ctx, constant.ERROR_LOAD_DATA, constant.FuncGetEmployee, "0010", "", "0000", err.Error())
			return nil, errors.New("data tidak ditemukan")
		}

		requestData = &response.GetDataArticleResponse{
			Message:    result.Message,
			StatusCode: result.StatusCode,
			Data:       result.Data,
		}
	}
	if claims.Role == "Admin" {
		resData, err := service.RequestEmployeeRepository.GetCompanyId(ctx, params.EmailUser)
		if err != nil {
			return nil, err
		}
		resDataAdmin = resData
		if resDataAdmin != "" {
			result, err := service.RequestEmployeeRepository.GetDataRequest(ctx, params.EmailUser, resDataAdmin)
			if err != nil {
				return nil, err

			}
			requestData = &response.GetDataArticleResponse{
				Message:    result.Message,
				StatusCode: result.StatusCode,
				Data:       result.Data,
			}

		}
	}

	return requestData, nil
}

func (service *requestArticleImpl) DeleteDataRequest(ctx *fiber.Ctx, params *request.GetDataRequest) (*fiber.Map, error) {
	if params.EmailUser == "" {
		message := "Email tidak boleh kosong"
		return nil, errors.New(message)
	}

	claims := utils.JwtVerificationWithClaim(ctx)
	if params.EmailUser != claims.Email {
		message := "Permintaan ditolak"
		return nil, errors.New(message)
	}
	if claims.Role != "Karyawan" {
		message := "Akses kamu untuk fitur ini ditolak"
		return nil, errors.New(message)
	}

	result, err := service.RequestEmployeeRepository.CheckIdRequest(params.Id)
	if err != nil {
		return nil, err
	}
	fmt.Println("result : ", result)
	if result > 0 {
		resultDel, err := service.RequestEmployeeRepository.DeleteDataRequest(params.Id)
		if err != nil {
			logging.Logging(ctx, constant.ERROR_DELETE_DATA, constant.FuncDeleteEmployee, "0010", "", "0000", err.Error())
			return nil, err
		}
		fmt.Println(resultDel)

		result := &fiber.Map{
			"Message": "Success",
		}
		return result, nil
	} else {
		result := &fiber.Map{
			"Message": "Id artikel tidak di temukan",
		}
		return result, nil
	}

}
