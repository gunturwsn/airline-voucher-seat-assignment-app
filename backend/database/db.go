package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// Init database connection for sqlite
func InitDB() {
	var err error

	db, err = sql.Open("sqlite3", "./vouchers.db")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(5)
	db.SetConnMaxIdleTime(5)
	createTable()
}

// Database migration for table voucher
func createTable() {
	createTabQuery := `
	CREATE TABLE IF NOT EXISTS vouchers (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        crew_name TEXT,
        crew_id TEXT,
        flight_number TEXT,
        flight_date TEXT,
        aircraft_type TEXT,
        seat1 TEXT,
        seat2 TEXT,
        seat3 TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
	`

	_, err := db.Exec(createTabQuery)
	if err != nil {
		log.Fatal(err)
	}
}

// Get connection to database
func GetDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./vouchers.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// Insert voucher data
func InsertVoucher(crewName, crewID, flightNumber, flightDate, aircraftType, seat1, seat2, seat3 string) error {
	db := GetDB()
	defer db.Close()

	query := `
    INSERT INTO vouchers (crew_name, crew_id, flight_number, flight_date, aircraft_type, seat1, seat2, seat3) 
    VALUES (?, ?, ?, ?, ?, ?, ?, ?)
    `
	_, err := db.Exec(query, crewName, crewID, flightNumber, flightDate, aircraftType, seat1, seat2, seat3)
	if err != nil {
		return err
	}
	return nil
}
