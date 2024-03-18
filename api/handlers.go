package api

import (
	"emp-mini/service"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {
	// Define API routes
	r.HandleFunc("/employee", service.CreateEmployee).Methods("POST")
	r.HandleFunc("/employee/{id}", service.UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employee/{id}", service.DeleteEmployee).Methods("DELETE")
	r.HandleFunc("/employee", service.GetAllEmployees).Methods("GET")
	r.HandleFunc("/employee/{id}", service.GetEmployeeByID).Methods("GET")

	r.HandleFunc("/rating", service.CreateRating).Methods("POST")
	r.HandleFunc("/rating/{id}", service.UpdateRating).Methods("PUT")
	r.HandleFunc("/rating/{id}", service.DeleteRating).Methods("DELETE")
	r.HandleFunc("/rating", service.GetAllRatings).Methods("GET")
	r.HandleFunc("/rating/{id}", service.GetRatingByID).Methods("GET")
	r.HandleFunc("/rating/{id}", service.PatchRating).Methods("PATCH")
}
