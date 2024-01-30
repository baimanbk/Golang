package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Employee struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Department string `json:"department,omitempty"`
	EmployeeID string `json:"employeeID,omitempty"`
	Age        int    `json:"age,omitempty"`
}

var employees []Employee

func main() {
	fmt.Println("Hello")
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/employee", createEmployee).Methods("POST")
	router.HandleFunc("/employee/{id}", getEmployeeByID).Methods("GET")
	router.HandleFunc("/employee/{id}", updateEmployee).Methods("PUT")
	router.HandleFunc("/employee/{id}", deleteEmployee).Methods("DELETE")
	router.HandleFunc("/employees", getAllEmployees).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

// Handler to create a new employee
func createEmployee(w http.ResponseWriter, r *http.Request) {
	var employee Employee
	_ = json.NewDecoder(r.Body).Decode(&employee)
	employee.ID = uuid.New().String()
	employees = append(employees, employee)
	json.NewEncoder(w).Encode(employee)
}

// Handler to get an employee by ID
func getEmployeeByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range employees {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Employee{})
}

// Handler to update an employee
func updateEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range employees {
		if item.ID == params["id"] {
			var updatedEmployee Employee
			_ = json.NewDecoder(r.Body).Decode(&updatedEmployee)
			updatedEmployee.ID = item.ID
			employees[index] = updatedEmployee
			json.NewEncoder(w).Encode(updatedEmployee)
			return
		}
	}
	json.NewEncoder(w).Encode(&Employee{})
}

// Handler to delete an employee by ID
func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range employees {
		if item.ID == params["id"] {
			employees = append(employees[:index], employees[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(employees)
}

// Handler to get all employees
func getAllEmployees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(employees)
}
