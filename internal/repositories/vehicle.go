package repositories

import (
	"github.com/jmoiron/sqlx"

	"github.com/Memchikkr/go-routes/internal/interfaces"
	"github.com/Memchikkr/go-routes/internal/models"
)

type VehicleRepository struct {
	db *sqlx.DB
}

func NewVehicleRepository(db *sqlx.DB) interfaces.VehicleRepository {
	return &VehicleRepository{db: db}
}

func (rep *VehicleRepository) GetVehicleById(vehicle_id int, user_id int) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	query := `SELECT * FROM "vehicle" WHERE id = $1 and user_id = $2`
	err := rep.db.Get(&vehicle, query, vehicle_id, user_id)

	if err != nil {
		return nil, err
	}

	return &vehicle, nil
}

func (rep *VehicleRepository) InsertVehicleData(user_id int, data *models.VehicleCreateRequest) (error) {
	query := `INSERT INTO "vehicle" (user_id, brand, license_plate)
				VALUES ($1, $2, $3)`
	_, err := rep.db.Exec(query, user_id, data.Brand, data.LicensePlate)
	return err
}

func (rep *VehicleRepository) GetVehiclesByUserId(user_id int) ([]*models.Vehicle, error) {
	var vehicle []*models.Vehicle
	query := `SELECT * FROM "vehicle" WHERE user_id = $1`
	err := rep.db.Select(&vehicle, query, user_id)
	if err != nil {
		return nil, err
	}
	return vehicle, nil
}

func (rep *VehicleRepository) DeleteVehicleById(vehicle_id int) (error) {
	query := `DELETE FROM "vehicle" WHERE id = $1`
	_, err := rep.db.Exec(query, vehicle_id)
	return err
}
