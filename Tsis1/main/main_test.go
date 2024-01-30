package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCRUD(t *testing.T) {
	employees = []Employee{
		{ID: "1", Name: "John Doe", Department: "IT", EmployeeID: "EMP-001", Age: 30},
		{ID: "2", Name: "Jane Smith", Department: "HR", EmployeeID: "EMP-002", Age: 35},
	}

	// Create
	employee := Employee{
		Name:       "John Doe",
		Department: "IT",
		EmployeeID: "EMP-003",
		Age:        30,
	}
	employeeJSON, _ := json.Marshal(employee)
	req, err := http.NewRequest("POST", "/employee", bytes.NewBuffer(employeeJSON))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createEmployee)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "Status code should be 200")

	// Read
	found := false
	for _, emp := range employees {
		if emp.Name == "John Doe" {
			found = true
			req, err = http.NewRequest("GET", "/employee/"+emp.ID, nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(getEmployeeByID)
			handler.ServeHTTP(rr, req)
			assert.Equal(t, http.StatusOK, rr.Code, "Status code should be 200")
			break
		}
	}
	if !found {
		t.Fatal("Employee not found")
	}

	// Update
	updatedEmployee := Employee{
		ID:         employee.ID,
		Name:       "Updated Name",
		Department: "Updated Department",
		EmployeeID: employee.EmployeeID,
		Age:        35,
	}
	updatedEmployeeJSON, _ := json.Marshal(updatedEmployee)
	req, err = http.NewRequest("PUT", "/employee/"+employee.ID, bytes.NewBuffer(updatedEmployeeJSON))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(updateEmployee)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "Status code should be 200")

	// Delete
	req, err = http.NewRequest("DELETE", "/employee/"+employee.ID, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(deleteEmployee)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "Status code should be 200")
}
