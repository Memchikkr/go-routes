package models

type Vehicle struct {
	ID           int    `json:"id" db:"id"`
	UserID       int    `json:"user_id" db:"user_id"`
	Brand        string `json:"brand" db:"brand"`
	LicensePlate string `json:"license_plate" db:"license_plate"`
}

type VehicleCreateRequest struct {
	Brand        string `json:"brand" db:"brand" example:"BMW"`
	LicensePlate string `json:"license_plate" db:"license_plate" example:"М798КМ136"`
}
