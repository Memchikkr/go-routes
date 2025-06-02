package interfaces

import "github.com/Memchikkr/go-routes/internal/models"

type TripRepository interface {
	InsertTripData(user_id int, data *models.TripCreateRequest) (error)
	GetTripById(trip_id int) (*models.Trip, error)
	GetTripByIdAndUserId(trip_id int, user_id int) (*models.Trip, error)
	GetAllTrips(filters *models.TripFilters) ([]*models.TripWithDetails, error)
	GetAllMyTrips(user_id int) ([]*models.TripWithDetails, error)
	UpdateTripComplete(trip_id int) (error)
	DeleteTripById(trip_id int) (error)
}

type TripService interface {
	InsertTripData(user_id int, data *models.TripCreateRequest) (error)
	GetTripById(trip_id int) (*models.Trip, error)
	GetAllTrips(filters *models.TripFilters) ([]*models.TripWithDetails, error)
	GetAllMyTrips(user_id int) ([]*models.TripWithDetails, error)
	UpdateTripComplete(trip_id int, user_id int) (error)
	DeleteTripById(trip_id int, user_id int) (error)
}
