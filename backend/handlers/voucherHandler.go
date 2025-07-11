package handlers

import (
	"backend/models"
	"backend/services"

	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

var voucherService = services.NewVoucherService()

// HandleError handles errors and returns them in a consistent format
func HandleError(w http.ResponseWriter, err error, statusCode int) {
	response := Response{
		Success: false,
		Message: err.Error(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// CheckVoucher handles requests to check whether a voucher already exists.
func CheckVoucher(w http.ResponseWriter, r *http.Request) {
	var input struct {
		FlightNumber string `json:"flightNumber"`
		Date         string `json:"date"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		HandleError(w, err, http.StatusBadRequest)
		return
	}

	exists, err := voucherService.CheckVoucherExists(input.FlightNumber, input.Date)
	if err != nil {
		HandleError(w, err, http.StatusInternalServerError)
		return
	}

	response := Response{
		Success: true,
		Data:    map[string]bool{"exists": exists},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GenerateVoucher handles requests to create new vouchers.
func GenerateVoucher(w http.ResponseWriter, r *http.Request) {
	var input models.Voucher
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		HandleError(w, err, http.StatusBadRequest)
		return
	}

	seats, err := voucherService.GenerateVoucher(input)
	if err != nil {
		HandleError(w, err, http.StatusInternalServerError)
		return
	}

	response := Response{
		Success: true,
		Data:    map[string]interface{}{"seats": seats},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
