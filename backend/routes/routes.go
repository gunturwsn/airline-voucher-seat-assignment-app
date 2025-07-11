package routes

import (
	"backend/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/api/check", handlers.CheckVoucher).Methods("POST")
	r.HandleFunc("/api/generate", handlers.GenerateVoucher).Methods("POST")
}
