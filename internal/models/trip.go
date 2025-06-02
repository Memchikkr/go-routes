package models

import (
	"time"

	"github.com/Memchikkr/go-routes/internal/utils"
)

type Trip struct {
	ID            int       `json:"id" db:"id"`
	DriverID      int       `json:"driver_id" db:"driver_id"`
	RouteID       int       `json:"route_id" db:"route_id"`
	VehicleID     int       `json:"vehicle_id" db:"vehicle_id"`
	DepartureTime time.Time `json:"departure_time" db:"departure_time"`
	ArrivalTime   time.Time `json:"arrival_time" db:"arrival_time"`
	SeatsCount    int       `json:"seats_count" db:"seats_count"`
	Price         float32   `json:"price" db:"price"`
	IsCompleted   bool      `json:"is_completed" db:"is_completed"`
}

type TripCreateRequest struct {
	RouteID       int       `json:"route_id" db:"route_id"`
	VehicleID     int       `json:"vehicle_id" db:"vehicle_id"`
	DepartureTime time.Time `json:"departure_time" db:"departure_time"`
	ArrivalTime   time.Time `json:"arrival_time" db:"arrival_time"`
	SeatsCount    int       `json:"seats_count" db:"seats_count"`
	Price         float32   `json:"price" db:"price"`
}

type TripFilters struct {
	From     string  `json:"from"`
	To       string  `json:"to"`
	Date     string  `json:"date"`
	MinSeats int     `json:"min_seats"`
	MaxPrice float64 `json:"max_price"`
}

type TripWithDetails struct {
	ID             int                     `db:"id" json:"id"`
	DriverID       int                     `json:"driver_id" db:"driver_id"`
	DepartureTime  time.Time               `db:"departure_time" json:"departure_time"`
	ArrivalTime    time.Time               `db:"arrival_time" json:"arrival_time"`
	SeatsCount     int                     `db:"seats_count" json:"seats_count"`
	Price          float64                 `db:"price" json:"price"`
	IsCompleted    bool                    `db:"is_completed" json:"is_completed"`
	DriverName     utils.NullToEmptyString `db:"name" json:"name"`
	DriverSurname  utils.NullToEmptyString `db:"surname" json:"surname"`
	DriverPhone    utils.NullToEmptyString `db:"phone_number" json:"phone_number"`
	DriverUsername utils.NullToEmptyString `db:"tg_username" json:"tg_username"`
	DriverDesc     utils.NullToEmptyString `db:"description" json:"description"`
	StartAddress   string                  `db:"start_address" json:"start_address"`
	StartLatitude  float64                 `json:"start_latitude" db:"start_latitude"`
	StartLongitude float64                 `json:"start_longitude" db:"start_longitude"`
	StopAddress    string                  `json:"stop_address" db:"stop_address"`
	StopLatitude   float64                 `json:"stop_latitude" db:"stop_latitude"`
	StopLongitude  float64                 `json:"stop_longitude" db:"stop_longitude"`
	Brand          string                  `db:"brand" json:"brand"`
	LicensePlate   string                  `db:"license_plate" json:"license_plate"`
	BookingsCount  int                     `db:"bookings_count" json:"bookings_count"`
	AvailableSeats int                     `db:"available_seats" json:"available_seats"`
}
