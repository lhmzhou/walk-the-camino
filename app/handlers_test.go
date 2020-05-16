package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
	"walk-the-camino/data"
	"walk-the-camino/database"

	"github.com/gorilla/mux"
)

type Employees struct {
	Emps []data.Employee
}

func Test_GetEmployee(t *testing.T) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/employees/{id}", getOneEmployee).Methods("GET")
	go func() {
		http.ListenAndServe("0.0.0.0:8787", router)
	}()
	time.Sleep(time.Second * 2)
	newEmployee := data.Employee{}
	newEmployee.ID = "1234"
	newEmployee.FirstName = "Higgs"
	newEmployee.LastName = "Boson"
	newEmployee.Company = "CERN"
	if ok := database.AddEmployee(newEmployee); !ok {
		t.Error("Error while adding employee to database")
		return
	}
	resp, err := http.Get("http://localhost:8787/employees/1234")
	if err != nil {
		t.Error("Sorry, unable to fetch employee")
		return
	}
	if resp.StatusCode != http.StatusOK {
		t.Error("Status not ok", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	respData := &data.Employee{}
	err = json.Unmarshal(body, respData)
	if err != nil {
		t.Error("Sorry, unable to unmarshall data", err.Error())
		return
	}
}

func Test_GetAllEmployee(t *testing.T) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/employees", getAllEmployee).Methods("GET")
	go func() {
		http.ListenAndServe("0.0.0.0:8788", router)
	}()
	time.Sleep(time.Second * 2)
	newEmployee := data.Employee{}
	newEmployee.ID = "1234"
	newEmployee.FirstName = "Higgs"
	newEmployee.LastName = "Boson"
	newEmployee.Company = "CERN"
	if ok := database.AddEmployee(newEmployee); !ok {
		t.Error("Error while adding employee to database")
		return
	}
	resp, err := http.Get("http://localhost:8788/employees")
	if err != nil {
		t.Error("Sorry, unable to fetch employee")
		return
	}
	if resp.StatusCode != http.StatusOK {
		t.Error("Status not ok", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	respData := []data.Employee{}
	err = json.Unmarshal(body, &respData)
	if err != nil {
		t.Error("Sorry, unable to unmarshall data", err.Error())
		return
	}
}

func Test_AddEmployee(t *testing.T) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/employee", createEmployee).Methods("POST")
	go func() {
		http.ListenAndServe("0.0.0.0:8789", router)
	}()
	time.Sleep(time.Second * 2)
	client := &http.Client{}
	url := fmt.Sprintf("http://%s:%s/%s", "localhost", "8789", "employee")
	emp := data.Employee{"1", "Higgs", "Boson", "Physicist", "Switzerland", "CERN"}
	oData, _ := json.Marshal(emp)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(oData))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Error("Error while sending transaction ", err)
		return
	}
	if resp.StatusCode != http.StatusCreated {
		t.Error("Unable to receive ok response", resp.StatusCode)
		return
	}

	emp = data.Employee{"1A", "Leonhard", "Euler", "Mathematician", "Basel", " University of Basel"}
	oData, _ = json.Marshal(emp)
	req, err = http.NewRequest("POST", url, bytes.NewBuffer(oData))
	req.Header.Set("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		t.Error("Error while sending transaction ", err)
		return
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Error("Unable receive ok response", resp.StatusCode)
		return
	}
	req, err = http.NewRequest("POST", url, bytes.NewBuffer([]byte("Hello")))
	req.Header.Set("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		t.Error("Error while sending transaction ", err)
		return
	}
	if resp.StatusCode != http.StatusBadRequest {
		t.Error("Unable to receive ok response", resp.StatusCode)
		return
	}
}

func Test_Delete_Employee(t *testing.T) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/employee/{id}", deleteEmployee).Methods("DELETE")
	go func() {
		http.ListenAndServe("0.0.0.0:8790", router)
	}()
	time.Sleep(time.Second * 2)
	emp = data.Employee{"1A", "Leonhard", "Euler", "Mathematician", "Basel", " University of Basel"}
	if ok := database.AddEmployee(newEmployee); !ok {
		t.Error("Error while adding employee to database")
		return
	}
	client := &http.Client{}
	url := fmt.Sprintf("http://%s:%s/%s", "localhost", "8790", "employee/1")
	req, err := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Error("Error while sending transaction ", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		t.Error("Unable to receive ok response", resp.StatusCode)
		return
	}
}

func Test_Update_Employee(t *testing.T) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/employees/{id}", updateEmpoyee).Methods("PUT")
	router.HandleFunc("/employees/{id}", getOneEmployee).Methods("GET")
	go func() {
		http.ListenAndServe("0.0.0.0:8791", router)
	}()
	time.Sleep(time.Second * 2)
	emp = data.Employee{"1A", "Leonhard", "Euler", "Mathematician", "Basel", " University of Basel"}
	if ok := database.AddEmployee(newEmployee); !ok {
		t.Error("Error while adding employee to database")
		return
	}
	client := &http.Client{}
	emp = data.Employee{"2A", "Ada", "Lovelace", "Mathematician", "London", " UK"}
	oData, _ := json.Marshal(emp)
	url := fmt.Sprintf("http://%s:%s/%s", "localhost", "8791", "employees/1")
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(oData))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Error("Error while sending transaction ", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		t.Error("Unable to receive ok response", resp.StatusCode)
		return
	}
	resp, err = client.Get(url)
	if err != nil {
		t.Error("Unable to get employee " + err.Error())
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &emp); err != nil {
		t.Error("Unable to unmarshal employee data " + err.Error())
	}
	if emp.Designation != "Mathematician" {
		t.Error("Failed to update employee")
	}
}

func Test_Flush_data(t *testing.T) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/flush", flushEmployee).Methods("GET")
	go func() {
		http.ListenAndServe("0.0.0.0:8792", router)
	}()
	time.Sleep(time.Second * 2)
	newEmployee := data.Employee{}
	newEmployee.ID = "5678"
	newEmployee.FirstName = "Katherine"
	newEmployee.LastName = "Johnson"
	newEmployee.Company = "NASA"
	if ok := database.AddEmployee(newEmployee); !ok {
		t.Error("Error while adding employee to database")
		return
	}
	newEmployee.ID = "1235"
	newEmployee.FirstName = "Vamsi"
	newEmployee.LastName = "Gorapalli"
	newEmployee.Company = "Intraedge"
	if ok := database.AddEmployee(newEmployee); !ok {
		t.Error("Error while adding employee to database")
		return
	}
	fmt.Println("Length of Employee Database ", len(database.GetEmployeeDatabase()))
	resp, err := http.Get("http://localhost:8792/flush")
	if err != nil {
		t.Error("unable to fetch employee")
		return
	}
	if resp.StatusCode != http.StatusOK {
		t.Error("status not ok", resp.StatusCode)
	}
	defer resp.Body.Close()
	if len(database.GetEmployeeDatabase()) != 0 {
		t.Errorf("Employee database length should be 0 but length is %d", len(database.GetEmployeeDatabase()))
	}
}
