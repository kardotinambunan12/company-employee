package repository

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"system_employee/config"
	errorhandler "system_employee/error_handler"
	"system_employee/model"
	"system_employee/model/request"
	"system_employee/model/response"
)

func NewRequestEmployeeUserRepository() RequestArticleRepository {
	return &requestEmployeeUserRepositoryImpl{}
}

type requestEmployeeUserRepositoryImpl struct{}

func (repository *requestEmployeeUserRepositoryImpl) InsertRequest(params *request.ArticleRequest) error {
	db := config.NewDB()

	sql := `insert into company_employee.article(title, descryption,company_id,email_user, user_created,created_at)values(?, ?, ?, ?,?,sysdate()) `
	result, err := db.Exec(sql, params.Title, params.Description, params.Data.EmployeeId, params.EmailUser, params.UserCreated)
	if err != nil {
		return err
	}
	defer db.Close()
	fmt.Println(result)

	return nil
}

func (repository *requestEmployeeUserRepositoryImpl) UpdateRequest(params *request.ArticleRequest) error {
	db := config.NewDB()
	fmt.Println("params : ", params)
	sql := `update company_employee.article set title=?, descryption=?,updated_at=sysdate() where id =?`
	result, err := db.Exec(sql, params.Title, params.Description, params.Id)
	if err != nil {
		return err
	}
	defer db.Close()
	fmt.Println(result)

	return nil
}

func (repository *requestEmployeeUserRepositoryImpl) GetDataRequest(ctx *fiber.Ctx, email string, restDataAdmin string) (*response.GetDataArticleResponse, error) {
	db := config.NewDB()
	employees := make([]response.EmployeeResponse, 0)
	fmt.Println("email : ", email)
	fmt.Println("restDataAdmin : ", restDataAdmin)

	if email != "" {
		sql := `select id, title, descryption, company_id, status, user_created from  company_employee.article where email_user=?`
		rows, err := db.Query(sql, email)
		if err != nil {
			errorhandler.PanicIfNeeded(err)
		}

		for rows.Next() {
			employee := response.EmployeeResponse{}
			err := rows.Scan(&employee.Id, &employee.Title, &employee.Description, &employee.EmployeeId, &employee.Status, &employee.UserCreated)
			if err != nil {
				errorhandler.PanicIfNeeded(err)
			}
			employees = append(employees, employee)
		}
	}
	if restDataAdmin != "" {

		sql := `select id, title, descryption, company_id, status, user_created from  company_employee.article where company_id=?`
		rows, err := db.Query(sql, restDataAdmin)
		if err != nil {
			errorhandler.PanicIfNeeded(err)
		}
		for rows.Next() {
			employee := response.EmployeeResponse{}
			err := rows.Scan(&employee.Id, &employee.Title, &employee.Description, &employee.EmployeeId, &employee.Status, &employee.UserCreated)
			if err != nil {
				errorhandler.PanicIfNeeded(err)
			}
			employees = append(employees, employee)
		}
	}

	data := &response.GetDataArticleResponse{
		StatusCode: "200",
		Message:    "Success",
		Data:       employees,
	}
	return data, nil
}

func (repository *requestEmployeeUserRepositoryImpl) DeleteDataRequest(id int) (*model.GeneralResponse, error) {
	db := config.NewDB()

	SQL := "delete from company_employee.article where id = ?"
	_, err := db.Exec(SQL, id)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	resData := &model.GeneralResponse{
		StatusCode: 200,
		Message:    "Success",
	}

	return resData, nil
}

func (repository *requestEmployeeUserRepositoryImpl) CheckIdRequest(id int) (int, error) {
	db := config.NewDB()

	sql := `select id from company_employee.article where id=? `

	rows, err := db.Query(sql, id)
	if err != nil {
		return 0, errors.New("Id tidak ditemukan")
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return 0, errors.New("id tidak di temukan")
		}
		return id, nil
	} else {
		return 0, errors.New("id tidak di temukan")
	}
}

func (repository *requestEmployeeUserRepositoryImpl) GetCompanyId(ctx *fiber.Ctx, email string) (string, error) {
	db := config.NewDB()
	var companyId string

	fmt.Println("id ", email)

	sql := `select company_id from company_employee.user where email=? `

	rows, err := db.Query(sql, email)
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
