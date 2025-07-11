package models

type Voucher struct {
	CrewName     string   `json:"name"`
	CrewID       string   `json:"id"`
	FlightNumber string   `json:"flightNumber"`
	Date         string   `json:"date"`
	AircraftType string   `json:"aircraft"`
	Seats        []string `json:"seats"`
}
