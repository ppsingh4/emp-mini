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

// HTTP Handler functions for Rating table
// POST
func CreateRating(w http.ResponseWriter, r *http.Request) {
	// Implement logic to create employee using HTTP

	decoder := json.NewDecoder(r.Body)
	var rate entity.Rating
	err := decoder.Decode(&rate)
	if err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	// Get the database instance
	dbInstance := db.GetDB()

	result := dbInstance.Create(&rate)
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
	var rate entity.Employee

	// Get the database instance
	dbInstance := db.GetDB()

	result := dbInstance.First(&rate, rateID)
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
	var ratings []entity.Rating

	// Get the database instance
	dbInstance := db.GetDB()

	result := dbInstance.Find(&ratings)
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

	// Get the database instance
	dbInstance := db.GetDB()

	// Find user by ID
	var rate entity.Rating
	result := dbInstance.First(&rate, id)
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
	dbInstance.Save(&rate)

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Rating of ID %d updated successfully", id)
}

// PATCH
func PatchRating(w http.ResponseWriter, r *http.Request) {
	// Extract rating ID from URL params
	vars := mux.Vars(r)
	ratingID := vars["id"]

	// Get the database instance
	dbInstance := db.GetDB()

	// Retrieve rating from the database
	var rating entity.Rating
	result := dbInstance.First(&rating, ratingID)
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
	result = dbInstance.Save(&rating)
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

	// Get the database instance
	dbInstance := db.GetDB()

	// Parse employee ID
	id, err := strconv.Atoi(rateID)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}
	// Delete employee record from the database
	result := dbInstance.Delete(&entity.Rating{}, rateID)
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
