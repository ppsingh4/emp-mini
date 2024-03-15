package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Model struct
type Employee struct {
	ID           uint
	First_Name   string
	Second_Name  string
	Email        string
	Phone_number string
	Department   string
}

type Rating struct {
	ID              uint
	Certification   string
	Task_Completion string
	Help            string
	EmployeeID      uint
}

var db *gorm.DB
var err error

func main() {
	// Connect to PostgreSQL Database
	time.Sleep(10 * time.Second)
	// db, err = gorm.Open(postgres.Open("postgresql://postgres:prithvi@db-service:5432/minidb?sslmode=disable"), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("failed to connect database")
	// }

	dsn := "host=postgres user=postgres password=prithvi dbname=minidb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// AutoMigrate the Employee schema
	err = db.AutoMigrate(&Employee{})
	if err != nil {
		log.Fatal("failed to migrate database")
	}

	// AutoMigrate the Rating schema
	err = db.AutoMigrate(&Rating{})
	if err != nil {
		log.Fatal("failed to migrate database")
	}

	// Initialize HTTP server with Gorilla Mux
	r := mux.NewRouter()
	// Define API routes
	r.HandleFunc("/employee", CreateEmployee).Methods("POST")
	r.HandleFunc("/employee/{id}", UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employee/{id}", DeleteEmployee).Methods("DELETE")
	r.HandleFunc("/employee", GetAllEmployees).Methods("GET")
	r.HandleFunc("/employee/{id}", GetEmployeeByID).Methods("GET")

	r.HandleFunc("/rating", CreateRating).Methods("POST")
	r.HandleFunc("/rating/{id}", UpdateRating).Methods("PUT")
	r.HandleFunc("/rating/{id}", DeleteRating).Methods("DELETE")
	r.HandleFunc("/rating", GetAllRatings).Methods("GET")
	r.HandleFunc("/rating/{id}", GetRatingByID).Methods("GET")
	r.HandleFunc("/rating/{id}", PatchRating).Methods("PATCH")

	// Start HTTP server
	fmt.Println("HTTP server listening on port 8080")
	http.ListenAndServe(":8080", r)
}

// ---------------------------------------------------------------------------
// HTTP Handler functions for Employee table
// POST
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	// Implement logic to create employee using HTTP

	decoder := json.NewDecoder(r.Body)
	var e Employee
	err := decoder.Decode(&e)
	if err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	result := db.Create(&e)
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
	var emp Employee
	result := db.First(&emp, empID)
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
	var employees []Employee
	result := db.Find(&employees)
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

	// Find user by ID
	var emp Employee
	result := db.First(&emp, id)
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
	db.Save(&emp)

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Employee %d updated successfully", id)
}

// DELETE
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	// Extract employee ID from the request URL
	vars := mux.Vars(r)
	empID := vars["id"]

	// Parse employee ID
	id, err := strconv.Atoi(empID)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}
	// Delete employee record from the database
	result := db.Delete(&Employee{}, empID)
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

//-------------------------------------------------------------------

// HTTP Handler functions for Rating table
// POST
func CreateRating(w http.ResponseWriter, r *http.Request) {
	// Implement logic to create employee using HTTP

	decoder := json.NewDecoder(r.Body)
	var rate Rating
	err := decoder.Decode(&rate)
	if err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	result := db.Create(&rate)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Rating for id %d created successfully!", rate.ID)
}

// GET
func GetRatingByID(w http.ResponseWriter, r *http.Request) {
	// Implement logic to retrieve user by ID using HTTP
	// Example:
	vars := mux.Vars(r)
	rateID := vars["id"]
	var rate Employee
	result := db.First(&rate, rateID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	// Return user details as JSON response
	json.NewEncoder(w).Encode(rate)
}

// GETALL
func GetAllRatings(w http.ResponseWriter, r *http.Request) {
	// Query all employee records from the database
	var ratings []Rating
	result := db.Find(&ratings)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Return employee details as JSON response
	json.NewEncoder(w).Encode(ratings)
}

// PUT
func UpdateRating(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rateID := vars["id"]

	// Parse user ID
	id, err := strconv.Atoi(rateID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Find user by ID
	var rate Rating
	result := db.First(&rate, id)
	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Decode JSON request body into User struct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&rate)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Update user record in the database
	db.Save(&rate)

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Rating of ID %d updated successfully", id)
}

// PATCH
func PatchRating(w http.ResponseWriter, r *http.Request) {
	// Extract rating ID from URL params
	vars := mux.Vars(r)
	ratingID := vars["id"]

	// Retrieve rating from the database
	var rating Rating
	result := db.First(&rating, ratingID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Decode JSON request body into a map[string]interface{}
	var updateData map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updateData); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Update rating fields based on the provided data
	for key, value := range updateData {
		switch key {
		case "certification":
			rating.Certification = value.(string)
		case "task_completion":
			rating.Task_Completion = value.(string)
		case "help":
			rating.Help = value.(string)
		}
	}

	// Save the updated rating to the database
	result = db.Save(&rating)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Rating %s updated successfully", ratingID)
}

// DELETE
func DeleteRating(w http.ResponseWriter, r *http.Request) {
	// Extract employee ID from the request URL
	vars := mux.Vars(r)
	rateID := vars["id"]

	// Parse employee ID
	id, err := strconv.Atoi(rateID)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}
	// Delete employee record from the database
	result := db.Delete(&Rating{}, rateID)
	if result.Error != nil {
		http.Error(w, "Failed to delete employee", http.StatusInternalServerError)
		return
	}
	// Check if any rows were affected
	if result.RowsAffected == 0 {
		http.Error(w, "Rating not found", http.StatusNotFound)
		return
	}
	// Respond with success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Rating of ID %d deleted successfully", id)
}
