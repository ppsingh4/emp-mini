package main

import (
	"emp-mini/api"
	"emp-mini/db"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//time.Sleep(10 * time.Second)
	// // Connect to PostgreSQL Database
	// if err := db.GetDB(); err != nil {
	// 	log.Fatal("Failed to initialize database:", err)
	// }

	// Initialize PostgreSQL Database
	db.InitDB()

	// Initialize HTTP server with Gorilla Mux
	r := mux.NewRouter()
	// Define API routes
	api.RegisterHandlers(r)

	// Start HTTP server

	http.ListenAndServe(":8080", r)
	fmt.Println("HTTP server listening on port 8080")

	// // Start the HTTP server
	// log.Println("Server listening on port 8080")
	// err := http.ListenAndServe(":8080", r)
	// if err != nil {
	// 	log.Fatal("Failed to start server:", err)
	// }
}
