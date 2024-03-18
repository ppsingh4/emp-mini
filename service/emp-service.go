package service

import (
	"emp-mini/db"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"emp-mini/entity"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	// Implement logic to create employee using HTTP

	decoder := json.NewDecoder(r.Body)
	var e entity.Employee
	err := decoder.Decode(&e)
	if err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	// Get the database instance
	dbInstance := db.GetDB()
	if dbInstance == nil {
		http.Error(w, "failed to obtain database instance", http.StatusInternalServerError)
		return
	}

	result := dbInstance.Create(&e)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Employee created successfully!")
}

// GET
func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	// Implement logic to retrieve user by ID using HTTP
	// Example:
	vars := mux.Vars(r)
	empID := vars["emp_id"]
	var emp entity.Employee

	// Get the database instance
	dbInstance := db.GetDB()

	result := dbInstance.First(&emp, empID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// Return user details as JSON response
	json.NewEncoder(w).Encode(emp)
}

// GETALL
func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	// Query all employee records from the database
	var employees []entity.Employee

	// Get the database instance
	dbInstance := db.GetDB()

	result := dbInstance.Find(&employees)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Return employee details as JSON response
	json.NewEncoder(w).Encode(employees)
}

// PUT
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empID := vars["id"]

	// Parse user ID
	id, err := strconv.Atoi(empID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Get the database instance
	dbInstance := db.GetDB()

	// Find user by ID
	var emp entity.Employee
	result := dbInstance.First(&emp, id)
	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Decode JSON request body into User struct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&emp)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Update user record in the database
	dbInstance.Save(&emp)

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Employee %d updated successfully", id)
}

// DELETE
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	// Extract employee ID from the request URL
	vars := mux.Vars(r)
	empID := vars["id"]

	// Get the database instance
	dbInstance := db.GetDB()

	// Parse employee ID
	id, err := strconv.Atoi(empID)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}
	// Delete employee record from the database
	result := dbInstance.Delete(&entity.Employee{}, empID)
	if result.Error != nil {
		http.Error(w, "Failed to delete employee", http.StatusInternalServerError)
		return
	}
	// Check if any rows were affected
	if result.RowsAffected == 0 {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	// Respond with success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Employee %d deleted successfully", id)
}
