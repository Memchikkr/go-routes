package services

import (
	"github.com/Memchikkr/go-routes/bootstrap"
	"github.com/Memchikkr/go-routes/internal/interfaces"
	"github.com/Memchikkr/go-routes/internal/models"
)

type TripService struct {
	repository interfaces.TripRepository
	env        *bootstrap.Env
}

func NewTripService(repository interfaces.TripRepository, env *bootstrap.Env) interfaces.TripService {
	return &TripService{repository: repository, env: env}
}

func (service *TripService) InsertTripData(user_id int, data *models.TripCreateRequest) error {
	err := service.repository.InsertTripData(user_id, data)
	return err
}

func (service *TripService) GetTripById(trip_id int) (*models.Trip, error) {
	trip, err := service.repository.GetTripById(trip_id)

	if err != nil {
		return nil, err
	}

	return trip, nil
}

func (service *TripService) GetAllTrips(filters *models.TripFilters) ([]*models.TripWithDetails, error) {
	trips, err := service.repository.GetAllTrips(filters)

	if err != nil {
		return nil, err
	}

	return trips, nil
}

func (service *TripService) GetAllMyTrips(user_id int) ([]*models.TripWithDetails, error) {
	trips, err := service.repository.GetAllMyTrips(user_id)

	if err != nil {
		return nil, err
	}

	return trips, nil
}

func (service *TripService) UpdateTripComplete(trip_id int, user_id int) error {
	_, err := service.repository.GetTripByIdAndUserId(trip_id, user_id)
	if err != nil {
		return err
	}
	upd_err := service.repository.UpdateTripComplete(trip_id)
	return upd_err
}

func (service *TripService) DeleteTripById(trip_id int, user_id int) error {
	_, err := service.repository.GetTripByIdAndUserId(trip_id, user_id)
	if err != nil {
		return err
	}
	del_err := service.repository.DeleteTripById(trip_id)
	return del_err
}
