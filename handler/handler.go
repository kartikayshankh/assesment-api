package handler

import (
	"api/endpoint"
	Service "api/service"

	"github.com/labstack/echo/v4"
)

func MakeHTTPHandler(service Service.HandlerInterface) *echo.Echo {
	e := echo.New()

	configEndpoint := endpoint.NewConfigEndpoint(service)

	e.POST("/CreateEmployee", configEndpoint.CreateEmployee)
	e.GET("/GetEmployeeByID/:id", configEndpoint.GetEmployeeByID)
	e.GET("/GetEmployee", configEndpoint.GetAllEmployee)
	e.PUT("/UpdateEmployee", configEndpoint.UpdateEmployeeByID)
	e.DELETE("/DeleteEmployee/:id", configEndpoint.DeleteEmployeeByID)
	return e
}
