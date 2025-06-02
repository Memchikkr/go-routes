package services

import (
	"github.com/Memchikkr/go-routes/bootstrap"
	"github.com/Memchikkr/go-routes/internal/interfaces"
	"github.com/Memchikkr/go-routes/internal/models"
)

type BookingService struct {
	repository interfaces.BookingRepository
	trip_repository interfaces.TripRepository
	env        *bootstrap.Env
}

func NewBookingService(repository interfaces.BookingRepository, trip_repository interfaces.TripRepository, env *bootstrap.Env) interfaces.BookingService {
	return &BookingService{repository: repository, trip_repository: trip_repository, env: env}
}

func (service *BookingService) InsertBookingData(trip_id int, user_id int) (error) {
	_, err := service.trip_repository.GetTripById(trip_id)
	if err != nil {
		return err
	}
	booking, err := service.repository.GetBookingRecordByTripIdAndUserId(trip_id, user_id)
	if booking != nil {
		return err
	}
	err_ins := service.repository.InsertBookingData(trip_id, user_id)
	return err_ins
}

func (service *BookingService) GetBookingRecordsByTripId(trip_id int) ([]*models.BookingWithDetails, error) {
	_, err := service.trip_repository.GetTripById(trip_id)
	if err != nil {
		return nil, err
	}
	bookings, err := service.repository.GetBookingRecordsByTripId(trip_id)

	if err != nil {
		return nil, err
	}

	return bookings, nil
}

func (service *BookingService) GetBookingRecordsByUserId(user_id int) ([]*models.BookingWithDetails, error) {
	bookings, err := service.repository.GetBookingRecordsByUserId(user_id)

	if err != nil {
		return nil, err
	}

	return bookings, nil
}

func (service *BookingService) DeleteBookingRecord(trip_id int, user_id int) (error) {
	_, err := service.trip_repository.GetTripById(trip_id)
	if err != nil {
		return err
	}
	booking, err := service.repository.GetBookingRecordByTripIdAndUserId(trip_id, user_id)
	if booking == nil {
		return err
	}
	err_del := service.repository.DeleteBookingRecord(trip_id, user_id)
	return err_del
}
