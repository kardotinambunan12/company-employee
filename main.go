package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"system_employee/config"
	"system_employee/controller"
	"system_employee/repository"
	"system_employee/service"
)

func main() {
	//configuration := config.New()

	// setup repository
	userEmployeeRepository := repository.NewCompanyEmployeeUserRepository()
	companyRepository := repository.NewCompanyEmployeerepository()
	requestEmployeeRespository := repository.NewRequestEmployeeUserRepository()

	// setup service
	userEmployeeService := service.NewCompanyEmployeeUserService(&userEmployeeRepository)
	companyService := service.NewCompanyEmployeeService(&companyRepository)
	requestCompanyService := service.NewRequestArticleService(&requestEmployeeRespository)

	//setup controller
	userEmployeeController := controller.NewCompanyEmployeeUserController(&userEmployeeService)
	companyController := controller.NewCompanyController(&companyService)
	requestCompanyController := controller.NewRequestemployeeController(&requestCompanyService)

	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	//app.Use(config.NewDB())

	app.Use(logger.New())

	// setup routing
	userEmployeeController.Route(app)
	companyController.Route(app)
	requestCompanyController.Route(app)

	err := app.Listen(":8080")
	if err != nil {
		return
	}

}
