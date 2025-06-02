package interfaces

import "github.com/Memchikkr/go-routes/internal/models"

type BookingRepository interface {
	InsertBookingData(trip_id int, user_id int) (error)
	GetBookingRecordsByTripId(trip_id int) ([]*models.BookingWithDetails, error)
	GetBookingRecordByTripIdAndUserId(trip_id int, user_id int) (*models.Booking, error)
	GetBookingRecordsByUserId(user_id int) ([]*models.BookingWithDetails, error)
	DeleteBookingRecord(trip_id int, user_id int) (error)
}

type BookingService interface {
	InsertBookingData(trip_id int, user_id int) (error)
	GetBookingRecordsByTripId(trip_id int) ([]*models.BookingWithDetails, error)
	GetBookingRecordsByUserId(user_id int) ([]*models.BookingWithDetails, error)
	DeleteBookingRecord(trip_id int, user_id int) (error)
}
