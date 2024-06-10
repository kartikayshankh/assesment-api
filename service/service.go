package service

import (
	"api/ent"
	"strconv"

	model "api/model"
	requestresponse "api/request_response"
	utils "api/utils"

	"github.com/labstack/echo/v4"
)

type HandlerInterface interface {
	CreateEmployee(e echo.Context, request *requestresponse.Employee) (*requestresponse.GenericResponse, *utils.ErrorHandler)
	GetEmployeeByID(e echo.Context, id string) (*requestresponse.Employee, *utils.ErrorHandler)
	GetAllEmployee(e echo.Context, limitString string, skipString string) (*requestresponse.EmployeeList, *utils.ErrorHandler)
	UpdateEmployeeByID(e echo.Context, request *requestresponse.Employee) (*requestresponse.GenericResponse, *utils.ErrorHandler)
	DeleteEmployeeByID(e echo.Context, id string) (*requestresponse.GenericResponse, *utils.ErrorHandler)
}

type Handler struct {
	db *ent.Client
}

// CreateEmployee implements HandlerInterface.

func NewHandler(db *ent.Client) HandlerInterface {
	return &Handler{db: db}
}

func (h *Handler) CreateEmployee(e echo.Context, request *requestresponse.Employee) (*requestresponse.GenericResponse, *utils.ErrorHandler) {
	response, err := model.CreateEmployee(h.db, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (h *Handler) GetEmployeeByID(e echo.Context, id string) (*requestresponse.Employee, *utils.ErrorHandler) {
	employee, err := model.GetEmployeeByID(h.db, id)
	if err != nil {
		// did error handling here utils.ErrorMapper(err.DevMessage)
		return nil, &utils.ErrorHandler{DevMessage: utils.ErrorMapper(err.DevMessage), Message: utils.ErrorMapper(err.DevMessage)}
	}
	if employee == nil {
		return nil, &utils.ErrorHandler{DevMessage: utils.DATA_NOT_FOUND, Message: utils.SOMETHING_WENT_WRONG}
	}
	return &requestresponse.Employee{
		ID:       employee.ID,
		Name:     employee.Name,
		Position: employee.Position,
		Salary:   employee.Salary,
	}, nil
}

func (h *Handler) GetAllEmployee(e echo.Context, limitString string, skipString string) (*requestresponse.EmployeeList, *utils.ErrorHandler) {
	employeeList := new(requestresponse.EmployeeList)
	limit, _ := strconv.Atoi(limitString)
	skip, _ := strconv.Atoi(skipString)
	skip = skip * limit
	response, err := model.GetAllEmployee(h.db, limit, skip)
	if err != nil {
		return nil, err
	}
	count := len(response)
	var employee requestresponse.Employee
	for _, value := range response {
		employee.ID = value.ID
		employee.Name = value.Name
		employee.Position = value.Position
		employee.Salary = value.Salary
		employeeList.Employee = append(employeeList.Employee, employee)
	}
	employeeList.Count = count
	return employeeList, nil
}

func (h *Handler) UpdateEmployeeByID(e echo.Context, request *requestresponse.Employee) (*requestresponse.GenericResponse, *utils.ErrorHandler) {

	/*employee, errEmployee := model.UpdateEmployeeByID(h.db, request)
	if errEmployee != nil {
		return nil, errEmployee
	}
	if employee == nil {
		return nil, &utils.ErrorHandler{DevMessage: utils.DATA_NOT_FOUND, Message: utils.SOMETHING_WENT_WRONG}
	}
	return &requestresponse.GenericResponse{
		Status:  utils.SUCCESS,
		Message: utils.EMPLOYEE_UPDATED_SUCCESSFULLY,
	}, nil */
	done := make(chan struct {
		employee *ent.Employee
		err      *utils.ErrorHandler
	})

	go func() {
		employee, errEmployee := model.UpdateEmployeeByID(h.db, request)
		done <- struct {
			employee *ent.Employee
			err      *utils.ErrorHandler
		}{employee, errEmployee}
	}()

	result := <-done
	if result.err != nil {
		return nil, result.err
	}
	if result.employee == nil {
		return nil, &utils.ErrorHandler{DevMessage: utils.DATA_NOT_FOUND, Message: utils.SOMETHING_WENT_WRONG}
	}
	return &requestresponse.GenericResponse{
		Status:  utils.SUCCESS,
		Message: utils.EMPLOYEE_UPDATED_SUCCESSFULLY,
	}, nil
}

func (h *Handler) DeleteEmployeeByID(e echo.Context, id string) (*requestresponse.GenericResponse, *utils.ErrorHandler) {
	errEmployee := model.DeleteEmployeeByID(h.db, id)
	if errEmployee != nil {
		return nil, errEmployee
	}
	return &requestresponse.GenericResponse{
		Status:  utils.SUCCESS,
		Message: utils.EMPLOYEE_DELETED_SUCCESSFULLY,
	}, nil
}
