package request

type ArticleRequest struct {
	Id          int          `json:"id"`
	Title       string       `json:"title"`
	EmailUser   string       `json:"emailUser"`
	Description string       `json:"description"`
	Data        DataEmployee `json:"data"`
	UserCreated string       `json:"userCreated"`
}

type GetDataRequest struct {
	Id        int    `json:"id"`
	EmailUser string `json:"emailUser"`
}

type ApprovalRequest struct {
	Id        int    `json:"id"`
	EmailUser string `json:"emailUser"`
	Status    string `json:"status"`
}

type DataEmployee struct {
	EmployeeId string `json:"employeeId"`
}

type InsertCompanyRequest struct {
	Title       string `json:"title"`
	CompanyId   string `json:"CompanyId"`
	Description string `json:"description"`
}
