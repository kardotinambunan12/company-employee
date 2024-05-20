package request

type UserRequest struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Address  string   `json:"address"`
	Data     DataType `json:"data"`
}

type DataType struct {
	RoleUser  string `json:"roleUser"`
	CompanyId string `json:"companyId"`
}

type GetEmailRequest struct {
	Email string `json:"email"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
