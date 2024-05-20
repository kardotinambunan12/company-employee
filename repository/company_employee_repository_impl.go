package repository

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"system_employee/config"
	errorhandler "system_employee/error_handler"
	"system_employee/model/request"
	"system_employee/model/response"
)

func NewCompanyEmployeerepository() CompanyEmployeeRepository {
	return &companyEmployeeRepositoryImpl{}
}

type companyEmployeeRepositoryImpl struct{}

func (repository *companyEmployeeRepositoryImpl) InserCompany(params *request.InsertCompanyRequest) error {

	db := config.NewDB()

	sql := `insert into company_employee.company(title, descryption,employee_id,created_at)values(?, ?, ?,sysdate()) `
	result, err := db.Exec(sql, params.Title, params.Description, params.CompanyId)
	if err != nil {
		return err
	}
	defer db.Close()
	fmt.Println(result)

	return nil

}

func (repository *companyEmployeeRepositoryImpl) ApprovalStatus(params *request.ApprovalRequest) (*response.ApprovalResponse, error) {

	db := config.NewDB()

	sql := `update company_employee.article set  status=? where id=? `
	result, err := db.Exec(sql, params.Status, params.Id)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	fmt.Println(result)
	dataApproval := response.DataApproval{
		Id:     params.Id,
		Status: params.Status,
	}
	rest := &response.ApprovalResponse{
		Message:    "success",
		StatusCode: "200",
		Data:       dataApproval,
	}

	return rest, nil

}

func (repository *companyEmployeeRepositoryImpl) CompanyList(ctx *fiber.Ctx) (*response.ListCompanyResponse, error) {
	db := config.NewDB()

	sql := `select id, title, descryption, employee_id from company_employee.company`
	rows, err := db.Query(sql)
	if err != nil {
		errorhandler.PanicIfNeeded(err)
	}

	companies := make([]response.EmployeeList, 0)
	for rows.Next() {
		company := response.EmployeeList{}
		err := rows.Scan(&company.Id, &company.Title, &company.Description, &company.CompanyId)
		if err != nil {
			errorhandler.PanicIfNeeded(err)
		}
		companies = append(companies, company)
	}

	data := &response.ListCompanyResponse{
		StatusCode: "200",
		Message:    "Success",
		Data:       companies,
	}
	return data, nil
}

func (repository *companyEmployeeRepositoryImpl) GetCompanyId(ctx *fiber.Ctx, id int) (string, error) {
	db := config.NewDB()
	var companyId string

	sql := `select company_id from company_employee.article where id=? `

	rows, err := db.Query(sql, id)
	if err != nil {
		return "", errors.New("Id tidak ditemukan")
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&companyId)
		if err != nil {
			return "", errors.New("id tidak di temukan")
		}
		return companyId, nil
	} else {
		return "", errors.New("id tidak di temukan")
	}
}
