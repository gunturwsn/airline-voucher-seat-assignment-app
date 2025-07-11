package services

import (
	"backend/database"
	"backend/interface/voucher"
	"backend/models"
	"fmt"
	"math/rand"
	"time"
)

type voucherService struct{}

// NewVoucherService creates a new instance of VoucherService
func NewVoucherService() voucher.VoucherService {
	return &voucherService{}
}

// CheckVoucherExists checks if the voucher already exists for a specific flight and date
func (v *voucherService) CheckVoucherExists(flightNumber, date string) (bool, error) {
	db := database.GetDB()

	defer db.Close()

	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM vouchers WHERE flight_number = ? AND flight_date = ?)"
	err := db.QueryRow(query, flightNumber, date).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// GenerateVoucher generates 3 random seats for the voucher and saves them in the database
func (v *voucherService) GenerateVoucher(data models.Voucher) ([]string, error) {
	// Generate seats based on the aircraft type
	seats := v.generateRandomSeats(data.AircraftType)

	// Save the voucher data into the database
	err := database.InsertVoucher(data.CrewName, data.CrewID, data.FlightNumber, data.Date, data.AircraftType, seats[0], seats[1], seats[2])
	if err != nil {
		return nil, err
	}
	database.GetDB().Close()

	return seats, nil
}

// GenerateRandomSeats generates 3 random seats according to the aircraft type
func (v *voucherService) generateRandomSeats(aircraftType string) []string {
	var seatList []string
	var rows int

	// Define seat layout and rows based on aircraft type
	switch aircraftType {
	case "ATR":
		rows = 18
	case "Airbus 320", "Boeing 737 Max":
		rows = 32
	default:
		return nil
	}

	// Seat letters based on aircraft type
	var seatLetters []string
	if aircraftType == "ATR" {
		seatLetters = []string{"A", "C", "D", "F"}
	} else {
		seatLetters = []string{"A", "B", "C", "D", "E", "F"}
	}

	rand.Seed(time.Now().UnixNano())

	// Generate 3 unique seats
	for len(seatList) < 3 {
		row := rand.Intn(rows) + 1
		col := seatLetters[rand.Intn(len(seatLetters))]
		seat := fmt.Sprintf("%d%s", row, col)

		// Check if the seat is already in the list
		if !contains(seatList, seat) {
			seatList = append(seatList, seat)
		}
	}

	return seatList
}

// contains checks if the seat already exists in the seat list
func contains(seats []string, seat string) bool {
	for _, s := range seats {
		if s == seat {
			return true
		}
	}
	return false
}
