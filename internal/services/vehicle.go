package services

import (
	"github.com/Memchikkr/go-routes/bootstrap"
	"github.com/Memchikkr/go-routes/internal/interfaces"
	"github.com/Memchikkr/go-routes/internal/models"
)

type VehicleService struct {
	repository interfaces.VehicleRepository
	env        *bootstrap.Env
}

func NewVehicleService(repository interfaces.VehicleRepository, env *bootstrap.Env) interfaces.VehicleService {
	return &VehicleService{repository: repository, env: env}
}

func (service *VehicleService) InsertVehicleData(user_id int, data *models.VehicleCreateRequest) (error) {
	err := service.repository.InsertVehicleData(user_id, data)
	return err
}

func (service *VehicleService) GetVehiclesByUserId(user_id int) ([]*models.Vehicle, error) {
	vehicles, err := service.repository.GetVehiclesByUserId(user_id)

	if err != nil {
		return nil, err
	}

	return vehicles, nil
}

func (service *VehicleService) DeleteVehicleById(vehicle_id int, user_id int) (error) {
	_, err := service.repository.GetVehicleById(vehicle_id, user_id)
	if err != nil {
		return err
	}
	del_err := service.repository.DeleteVehicleById(vehicle_id)
	return del_err
}
