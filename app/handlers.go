package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"walk-the-camino/data"
	"walk-the-camino/database"

	"github.com/gorilla/mux"
)

func createEmployee(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received to create employee")
	var newEmployee data.Employee
	// Convert r.Body into a readable formart
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the employee id, title and description only in order to update")
	}

	if err := json.Unmarshal(reqBody, &newEmployee); err != nil {
		log.Println("Unable to unmarshal incoming post body " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to unmarshal incoming post body"))
		return
	}
	if ok := database.AddEmployee(newEmployee); !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Return the 201 created status code
	w.WriteHeader(http.StatusCreated)
	// Return the newly created employee
	json.NewEncoder(w).Encode(newEmployee)
}

func getOneEmployee(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the url
	empId := mux.Vars(r)["id"]
	log.Printf("Request received to get employee id [%s]\n", empId)
	// Get the details from an existing employee
	// Use the blank identifier to avoid creating a value that will not be used
	for _, singleEvent := range database.GetEmployees() {
		if singleEvent.ID == empId {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func getAllEmployee(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received to get all employees")
	var employees = []data.Employee{}
	for _, emp := range database.GetEmployees() {
		employees = append(employees, emp)
	}
	json.NewEncoder(w).Encode(employees)
}

func updateEmpoyee(w http.ResponseWriter, r *http.Request) {
	employeeID := mux.Vars(r)["id"]
	log.Printf("Request received to update employee id [%s]\n", employeeID)
	var updatedEmployee data.Employee
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the employee title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedEmployee)
	for _, employee := range database.GetEmployees() {
		if employee.ID == employeeID {
			employee.FirstName = updatedEmployee.FirstName
			employee.LastName = updatedEmployee.LastName
			employee.Company = updatedEmployee.Company
			employee.Designation = updatedEmployee.Designation
			database.UpdateEmployee(employee)
			json.NewEncoder(w).Encode(employee)
		}
	}
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	//Get the ID from the url
	empID := mux.Vars(r)["id"]
	log.Printf("Request received to delete employee id [%s]\n", empID)
	// Get the details from an existing employee
	// Use the blank identifier to avoid creating a value that will not be used
	for _, employee := range database.GetEmployees() {
		if employee.ID == empID {
			//delete(employeeDatabase, employee.ID)
			database.DeleteEmployee(empID)
			fmt.Fprintf(w, "The employee with ID %v has been deleted successfully", empID)
		}
	}
}

func flushEmployee(w http.ResponseWriter, r *http.Request) {
	//Get the ID from the url
	log.Printf("Request received to flush all data\n")
	// Get the details from an existing employee
	// Use the blank identifier to avoid creating a value that will not be used
	database.FlushEmployees()
	w.WriteHeader(http.StatusOK)
	return
}
