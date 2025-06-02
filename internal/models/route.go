package models

type Route struct {
	ID             int     `json:"id" db:"id"`
	UserID         int     `json:"user_id" db:"user_id"`
	StartAddress   string  `json:"start_address" db:"start_address"`
	StartLatitude  float64 `json:"start_latitude" db:"start_latitude"`
	StartLongitude float64 `json:"start_longitude" db:"start_longitude"`
	StopAddress    string  `json:"stop_address" db:"stop_address"`
	StopLatitude   float64 `json:"stop_latitude" db:"stop_latitude"`
	StopLongitude  float64 `json:"stop_longitude" db:"stop_longitude"`
}

type RouteCreateRequest struct {
	StartAddress   string  `json:"start_address" db:"start_address" example:"Moscow, Russia"`
	StartLatitude  float64 `json:"start_latitude" db:"start_latitude" example:"42.223"`
	StartLongitude float64 `json:"start_longitude" db:"start_longitude" example:"42.223"`
	StopAddress    string  `json:"stop_address" db:"stop_address" example:"Moscow, Russia"`
	StopLatitude   float64 `json:"stop_latitude" db:"stop_latitude" example:"42.223"`
	StopLongitude  float64 `json:"stop_longitude" db:"stop_longitude" example:"42.223"`
}
