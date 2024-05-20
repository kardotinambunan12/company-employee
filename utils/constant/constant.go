package constant

const (
	APP_NAME = "SYSTEM_EMPLOYEE"
)

const (
	ERRORMIDDLEWARE   = "ERROR MIDDLEWARE"
	ERROR_LOAD_DATA   = "GAGAL MENGAMBIL DATA DARI DATABASE"
	ERROR_UPDATE_DATA = "GAGAL UPDATE DATA KE DATABASE"
	ERROR_INSERT_DATA = "GAGAL INSERT DATA KE DATABASE"
	ERROR_DELETE_DATA = "GAGAL DELETE DATA KE DATABASE"
)

//service

const (
	FuncInsertEmployee = "companyEmployeeImpl.InsertEmployee"
	FuncApproval       = "companyEmployeeImpl.AdminApproval"
	FuncDeleteEmployee = "requestEmployeeImpl.DeleteDataRequest"
	FuncGetEmployee    = "requestEmployeeImpl.GetDataRequest"
	FuncUpdateEmployee = "requestEmployeeImpl.UpdateRequest"
	//FuncInsertEmployee = "requestEmployeeImpl.InsertRequest"
)
