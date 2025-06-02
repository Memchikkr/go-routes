package interfaces

import "github.com/Memchikkr/go-routes/internal/models"

type VehicleRepository interface {
	InsertVehicleData(user_id int, data *models.VehicleCreateRequest) (error)
	GetVehiclesByUserId(user_id int) ([]*models.Vehicle, error)
	GetVehicleById(vehicle_id int, user_id int) (*models.Vehicle, error)
	DeleteVehicleById(vehicle_id int) (error)
}

type VehicleService interface {
	InsertVehicleData(user_id int, data *models.VehicleCreateRequest) (error)
	GetVehiclesByUserId(user_id int) ([]*models.Vehicle, error)
	DeleteVehicleById(vehicle_id int, user_id int) (error)
}
