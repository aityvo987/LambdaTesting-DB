package usecase

import (
	entity "lambda-test/internal/entity"
	"lambda-test/internal/repository"
	model "lambda-test/internal/repository/model"
	"log"

	"github.com/jinzhu/copier"
)

type EmployeeUsecase interface {
	GetEmployees() ([]model.Employee, error)
	GetEmployee() ([]model.Employee, error)
	CreateEmployee(employee model.Employee) error
	UpdateEmployee(id int, employee model.Employee) error
	DeleteEmployee(id int) error
}

func GetEmployees(req entity.GetEmployeesRequest) (entity.GetEmployeesResponse, error) {
	resp := entity.GetEmployeesResponse{}
	log.Printf("Read Employees - Calling to Repository\n")
	employees, err := repository.ReadAllEmployee()
	if err != nil {
		return resp, err
	}
	resp.Employees = employees
	return resp, nil
}

func GetEmployee(req entity.GetEmployeeRequest) (entity.GetEmployeeResponse, error) {
	resp := entity.GetEmployeeResponse{}
	log.Printf("Read Employee - Calling to Repository\n")
	employee, err := repository.ReadEmployee(req.EmployeeID)
	if err != nil {
		return resp, err
	}
	resp.Employee = employee
	return resp, nil
}

func CreateEmployee(req entity.CreateEmployeeRequest) error {
	emp := model.Employee{}
	copier.Copy(&emp, &req.Data)
	log.Printf("Create Employee - Calling to Repository\n")
	err := repository.CreateEmployee(emp)
	if err != nil {
		return err
	}
	return nil
}

func UpdateEmployee(req entity.UpdateEmployeeRequest) error {
	emp := model.Employee{}
	copier.Copy(&emp, &req.Data)
	log.Printf("Update Employee - Calling to Repository\n")
	err := repository.UpdateEmployee(req.EmployeeID, emp)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEmployee(req entity.DeleteEmployeeRequest) error {
	log.Printf("Delete Employee - Calling to Repository\n")
	err := repository.DeleteEmployee(req.EmployeeID)
	if err != nil {
		return err
	}
	return nil
}
