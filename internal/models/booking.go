package models

import (
	"time"

	"github.com/Memchikkr/go-routes/internal/utils"
)

type Booking struct {
	TripID      int       `json:"trip_id" db:"trip_id"`
	UserID      int       `json:"user_id" db:"user_id"`
	BookingTime time.Time `json:"booking_time" db:"booking_time"`
	IsApproved  bool      `json:"is_approved" db:"is_approved"`
}

type BookingWithDetails struct {
	TripID         int                     `json:"trip_id" db:"trip_id"`
	UserID		   int                     `json:"user_id" db:"user_id"`
	BookingTime    time.Time               `json:"booking_time" db:"booking_time"`
	IsApproved     bool                    `json:"is_approved" db:"is_approved"`
	DepartureTime  time.Time               `db:"departure_time" json:"departure_time"`
	ArrivalTime    time.Time               `db:"arrival_time" json:"arrival_time"`
	Price          float64                 `db:"price" json:"price"`
	IsCompleted    bool                    `db:"is_completed" json:"is_completed"`
	DriverID       int                     `json:"driver_id" db:"driver_id"`
	DriverUsername utils.NullToEmptyString `db:"tg_username" json:"tg_username"`
	DriverPhone    utils.NullToEmptyString `db:"phone_number" json:"phone_number"`
	StartAddress   string                  `db:"start_address" json:"start_address"`
	StartLatitude  float64                 `json:"start_latitude" db:"start_latitude"`
	StartLongitude float64                 `json:"start_longitude" db:"start_longitude"`
	StopAddress    string                  `json:"stop_address" db:"stop_address"`
	StopLatitude   float64                 `json:"stop_latitude" db:"stop_latitude"`
	StopLongitude  float64                 `json:"stop_longitude" db:"stop_longitude"`
	Brand          string                  `db:"brand" json:"brand"`
	LicensePlate   string                  `db:"license_plate" json:"license_plate"`
}
