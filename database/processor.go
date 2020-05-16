package database

import (
	"walk-the-camino/data"
	"strconv"
	"time"
)

var employeeDatabase map[string]data.Employee

func init() {
	employeeDatabase = make(map[string]data.Employee)
}

func AddEmployee(employee data.Employee) bool {
	if _, err := strconv.Atoi(employee.ID); err != nil {
		return false
	}
	employeeDatabase[employee.ID] = employee
	return true
}

func DeleteEmployee(id string) bool {
	if _, ok := employeeDatabase[id]; !ok {
		return false
	}
	delete(employeeDatabase,id)
	return true
}

func GetEmployees() map[string]data.Employee {
	return employeeDatabase
}

func UpdateEmployee(employee data.Employee) bool {
	time.Sleep(2*time.Second)
	employeeDatabase[employee.ID] = employee
	return true
}

func FlushEmployees() {
	for k := range employeeDatabase {
		delete(employeeDatabase, k)
	}
}

func GetEmployeeDatabase() map[string]data.Employee {
	return employeeDatabase
}
