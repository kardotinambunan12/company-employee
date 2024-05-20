package repository

import (
	"errors"
	"fmt"
	"system_employee/config"
	"system_employee/model/request"
	"system_employee/model/response"
)

func NewCompanyEmployeeUserRepository() CompanyEmployeeUserRepository {
	return &companyEmployeeUserRepositoryImpl{}
}

type companyEmployeeUserRepositoryImpl struct{}

func (r *companyEmployeeUserRepositoryImpl) GetEmail(request *request.GetEmailRequest) *response.EmailResponse {
	db := config.NewDB()

	sql := `SELECT email FROM company_employee.user WHERE email = ?`
	fmt.Println("Querying email: ", request.Email)
	row := db.QueryRow(sql, request.Email)

	resultEmail := &response.EmailResponse{}
	err := row.Scan(&resultEmail.Email)
	if err != nil {
		// Handle case when email is not found or other errors
		return &response.EmailResponse{Email: ""}
	}

	// Check if email is empty in database
	if resultEmail.Email == "" {
		return &response.EmailResponse{Email: ""}
	}

	return resultEmail

}

func (r *companyEmployeeUserRepositoryImpl) UserRegister(params *request.UserRequest) error {
	db := config.NewDB()
	sql := `INSERT INTO company_employee.user
    							(name, password, address, email, role, company_id, created_at)
				VALUES (?, ?, ?, ?, ?, ?, SYSDATE())`
	result, err := db.Exec(sql, params.Name, params.Password, params.Address, params.Email, params.Data.RoleUser, params.Data.CompanyId)
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}

func (r *companyEmployeeUserRepositoryImpl) Login(params *request.LoginRequest) (*response.LoginResponse, error) {
	db := config.NewDB()
	sql := `select email from company_employee.user where email = ? and password = ?`

	rows, _ := db.Query(sql, params.Email, params.Password)

	defer rows.Close()

	resultLogin := &response.LoginResponse{}
	if rows.Next() {
		err := rows.Scan(&resultLogin.Email)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("data register not found")
	}
	rsLogin := &response.LoginResponse{
		Message:      "success",
		IsSuccessful: true,
		Email:        resultLogin.Email,
	}

	return rsLogin, nil
}
