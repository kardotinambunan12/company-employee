package response

type EmailResponse struct {
	Email string `json:"email"`
}

type LoginResponse struct {
	IsSuccessful bool   `json:"isSuccessful"`
	Email        string `json:"email,omitempty"`
	Message      string `json:"message"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RoleResponse struct {
	RoleUser string `json:"roleUser"`
	Email    string `json:"email"`
}

type EmployeeResponse struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	EmployeeId  string `json:"employeeId"`
	Status      string `json:"status"`
	UserCreated string `json:"userCreated"`
}

type GetDataArticleResponse struct {
	Message    string             `json:"message"`
	StatusCode string             `json:"statusCode"`
	Data       []EmployeeResponse `json:"data"`
}

type ApprovalResponse struct {
	Message    string       `json:"message"`
	StatusCode string       `json:"statusCode"`
	Data       DataApproval `json:"data"`
}
type DataApproval struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

type ListCompanyResponse struct {
	Message    string         `json:"message"`
	StatusCode string         `json:"statusCode"`
	Data       []EmployeeList `json:"data"`
}

type EmployeeList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CompanyId   string `json:"CompanyId"`
}
