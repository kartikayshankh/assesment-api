package endpoint

import (
	requestresponse "api/request_response"
	Service "api/service"
	"api/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ConfigEndpoint interface {
	CreateEmployee(e echo.Context) error
	GetEmployeeByID(e echo.Context) error
	GetAllEmployee(e echo.Context) error
	UpdateEmployeeByID(e echo.Context) error
	DeleteEmployeeByID(e echo.Context) error
}

type configEndpoint struct {
	service Service.HandlerInterface
}

func NewConfigEndpoint(service Service.HandlerInterface) ConfigEndpoint {
	return &configEndpoint{service: service}
}

func (c *configEndpoint) CreateEmployee(e echo.Context) error {
	request := new(requestresponse.Employee)

	err := e.Bind(&request)
	if err != nil {
		return e.JSON(http.StatusBadRequest, &utils.ErrorHandler{DevMessage: err.Error(), Message: utils.SOMETHING_WENT_WRONG})
	}
	if request.Name == "" || request.Position == "" || request.Salary == 0.0 {
		return e.JSON(http.StatusBadRequest, &utils.ErrorHandler{DevMessage: utils.DATA_MISSING, Message: utils.SOMETHING_WENT_WRONG})
	}
	response, errResponse := c.service.CreateEmployee(e, request)
	if errResponse != nil {
		return e.JSON(http.StatusBadRequest, errResponse)
	}
	return e.JSON(http.StatusCreated, response)
}

func (c *configEndpoint) GetEmployeeByID(e echo.Context) error {
	id := e.Param("id")
	if id == "" {
		return e.JSON(http.StatusBadRequest, &utils.ErrorHandler{DevMessage: utils.DATA_MISSING, Message: utils.SOMETHING_WENT_WRONG})
	}
	response, errResponse := c.service.GetEmployeeByID(e, id)
	if errResponse != nil {
		return e.JSON(http.StatusBadRequest, errResponse)
	}
	return e.JSON(http.StatusOK, response)
}

func (c *configEndpoint) GetAllEmployee(e echo.Context) error {
	skip := e.QueryParam("skip")
	limit := e.QueryParam("limit")
	response, errResponse := c.service.GetAllEmployee(e, limit, skip)
	if errResponse != nil {
		return e.JSON(http.StatusNoContent, errResponse)
	}
	return e.JSON(http.StatusOK, response)
}

func (c *configEndpoint) UpdateEmployeeByID(e echo.Context) error {
	request := new(requestresponse.Employee)

	err := e.Bind(&request)
	if err != nil {
		return e.JSON(http.StatusBadRequest, &utils.ErrorHandler{DevMessage: err.Error(), Message: utils.SOMETHING_WENT_WRONG})
	}
	if request.ID == 0 {
		return e.JSON(http.StatusBadRequest, &utils.ErrorHandler{DevMessage: utils.DATA_MISSING, Message: utils.SOMETHING_WENT_WRONG})
	}
	response, errResponse := c.service.UpdateEmployeeByID(e, request)
	if errResponse != nil {
		return e.JSON(http.StatusBadRequest, errResponse)
	}
	return e.JSON(http.StatusOK, response)
}

func (c *configEndpoint) DeleteEmployeeByID(e echo.Context) error {
	id := e.Param("id")
	if id == "" {
		return e.JSON(http.StatusBadRequest, &utils.ErrorHandler{DevMessage: utils.DATA_MISSING, Message: utils.SOMETHING_WENT_WRONG})
	}
	response, errResponse := c.service.DeleteEmployeeByID(e, id)
	if errResponse != nil {
		return e.JSON(http.StatusNoContent, errResponse)
	}
	return e.JSON(http.StatusOK, response)
}
