package model

import (
	"api/ent"
	"api/ent/employee"
	"api/ent/predicate"
	requestresponse "api/request_response"
	"api/utils"
	"context"
	"strconv"
)

func CreateEmployee(db *ent.Client, request *requestresponse.Employee) (*requestresponse.GenericResponse, *utils.ErrorHandler) {
	newclient := db.Employee.Create()

	newclient.SetName(request.Name)
	newclient.SetPosition(request.Position)
	newclient.SetSalary(request.Salary)

	employee, errEmployee := newclient.Save(context.Background())
	if errEmployee != nil {
		return nil, &utils.ErrorHandler{DevMessage: errEmployee.Error(), Message: errEmployee.Error()}
	}
	if employee != nil {
		return &requestresponse.GenericResponse{
			Status:  utils.SUCCESS,
			Message: utils.EMPLOYEE_CREATED_SUCCESSFULLY,
		}, nil
	}
	return &requestresponse.GenericResponse{
		Status:  utils.FAILED,
		Message: utils.SOMETHING_WENT_WRONG,
	}, nil
}

func GetEmployeeByID(db *ent.Client, id string) (*ent.Employee, *utils.ErrorHandler) {
	condition := employeeQueryGenerator(id)
	employee, errEmployee := db.Employee.Query().
		Where(condition).
		Only(context.Background())
	if errEmployee != nil {
		return nil, &utils.ErrorHandler{DevMessage: errEmployee.Error(), Message: errEmployee.Error()}
	}
	return employee, nil
}

func GetAllEmployee(db *ent.Client, limit, skip int) ([]*ent.Employee, *utils.ErrorHandler) {
	employee, errEmployee := db.Employee.Query().
		Offset(skip).
		Limit(limit).
		All(context.Background())
	if errEmployee != nil {
		return nil, &utils.ErrorHandler{DevMessage: errEmployee.Error(), Message: errEmployee.Error()}
	}
	return employee, nil
}

func UpdateEmployeeByID(db *ent.Client, request *requestresponse.Employee) (*ent.Employee, *utils.ErrorHandler) {
	fetchEmployee, errFetchEmployee := db.Employee.Query().
		Where(employee.IDEQ(request.ID)).
		Only(context.Background())

	if errFetchEmployee != nil {
		return nil, &utils.ErrorHandler{DevMessage: errFetchEmployee.Error(), Message: errFetchEmployee.Error()}
	}
	update := fetchEmployee.Update()

	if request.Name != "" {
		update.SetName(request.Name)
	}
	if request.Position != "" {
		update.SetPosition(request.Position)
	}
	if request.Salary > 0.0 {
		update.SetSalary(request.Salary)
	}
	updateEmployee, errUpdate := update.Save(context.Background())
	if errUpdate != nil {
		return nil, &utils.ErrorHandler{DevMessage: errUpdate.Error(), Message: errUpdate.Error()}
	}
	return updateEmployee, nil
}

func DeleteEmployeeByID(db *ent.Client, id string) *utils.ErrorHandler {
	// condition := employeeQueryGenerator(id)
	idInt, _ := strconv.Atoi(id)
	responseId := db.Employee.DeleteOneID(idInt).
		Exec(context.Background())
	if responseId != nil {
		return &utils.ErrorHandler{DevMessage: responseId.Error(), Message: responseId.Error()}
	}
	return nil
}

func employeeQueryGenerator(id string) predicate.Employee {
	idInt, _ := strconv.Atoi(id)
	var condition predicate.Employee
	if id != "" {
		condition = employee.ID(idInt)
	}
	return condition
}
