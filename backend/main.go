package main

import (
	"backend/database"
	"backend/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Init database connection
	database.InitDB()

	// Setup router and route API
	r := mux.NewRouter()

	// Setup routes for voucher checking and voucher creation
	r.HandleFunc("/api/check", handlers.CheckVoucher).Methods("POST")
	r.HandleFunc("/api/generate", handlers.GenerateVoucher).Methods("POST")

	// Setup CORS to allow frontend from localhost:3000
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// Wrap router with CORS handler
	handler := c.Handler(r)

	// Server running on port 8000
	fmt.Println("Server running on port http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
