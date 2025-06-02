package repositories

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/Memchikkr/go-routes/internal/interfaces"
	"github.com/Memchikkr/go-routes/internal/models"
)

type TripRepository struct {
	db *sqlx.DB
}

func NewTripRepository(db *sqlx.DB) interfaces.TripRepository {
	return &TripRepository{db: db}
}

func (rep *TripRepository) GetTripByIdAndUserId(trip_id int, user_id int) (*models.Trip, error) {
	var trip models.Trip
	query := `SELECT * FROM "trip" WHERE id = $1 and driver_id = $2`
	err := rep.db.Get(&trip, query, trip_id, user_id)

	if err != nil {
		return nil, err
	}

	return &trip, nil
}

func (rep *TripRepository) InsertTripData(user_id int, data *models.TripCreateRequest) error {
	query := `INSERT INTO "trip" (driver_id, route_id, vehicle_id, departure_time, arrival_time, seats_count, price)
				VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := rep.db.Exec(query, user_id, &data.RouteID, &data.VehicleID, &data.DepartureTime, &data.ArrivalTime, &data.SeatsCount, &data.Price)
	return err
}

func (rep *TripRepository) GetTripById(trip_id int) (*models.Trip, error) {
	var trip models.Trip
	query := `SELECT * FROM "trip" WHERE id = $1`
	err := rep.db.Get(&trip, query, trip_id)

	if err != nil {
		return nil, err
	}

	return &trip, nil
}


func (rep *TripRepository) GetAllTrips(filters *models.TripFilters) ([]*models.TripWithDetails, error) {
	query := `
		SELECT 
			t.id, 
			t.driver_id,
			t.departure_time, 
			t.arrival_time, 
			t.seats_count,
			t.price, 
			t.is_completed,
			u.name, 
			u.surname,
			u.phone_number,
			u.tg_username, 
			u.description,
			r.start_address,
			r.start_latitude,
			r.start_longitude,
			r.stop_address,
			r.stop_latitude,
			r.stop_longitude,
			v.brand,
			v.license_plate,
			(
				SELECT COUNT(*) 
				FROM booking b 
				WHERE b.trip_id = t.id AND b.is_approved = true
			) AS bookings_count,
			t.seats_count - (
				SELECT COUNT(*) 
				FROM booking b 
				WHERE b.trip_id = t.id AND b.is_approved = true
			) AS available_seats
		FROM 
			trip t
		JOIN 
			"user" u ON t.driver_id = u.id
		JOIN 
			route r ON t.route_id = r.id
		JOIN 
			vehicle v ON t.vehicle_id = v.id
		LEFT JOIN 
			booking b ON t.id = b.trip_id
		WHERE 
			NOT t.is_completed
	`
	args := []interface{}{}
	argPos := 1

	if filters.From != "" {
		query += fmt.Sprintf(" AND r.start_address LIKE $%d", argPos)
		args = append(args, "%"+filters.From+"%")
		argPos++
	}

	if filters.To != "" {
		query += fmt.Sprintf(" AND r.stop_address LIKE $%d", argPos)
		args = append(args, "%"+filters.To+"%")
		argPos++
	}

	if filters.Date != "" {
		query += fmt.Sprintf(" AND DATE(t.departure_time) < $%d", argPos)
		args = append(args, filters.Date)
		argPos++
	}

	if filters.MinSeats > 0 {
		query += fmt.Sprintf(" AND t.seats_count >= $%d", argPos)
		args = append(args, filters.MinSeats)
		argPos++
	}

	if filters.MaxPrice > 0 {
		query += fmt.Sprintf(" AND t.price <= $%d", argPos)
		args = append(args, filters.MaxPrice)
		argPos++
	}

	query += " ORDER BY t.departure_time ASC"

	var trips []*models.TripWithDetails
	err := rep.db.Select(&trips, query, args...)
	return trips, err
}

func (rep *TripRepository) GetAllMyTrips(user_id int) ([]*models.TripWithDetails, error) {
	var trips []*models.TripWithDetails
	query := `
		SELECT 
			t.id, 
			t.driver_id,
			t.departure_time, 
			t.arrival_time, 
			t.seats_count,
			t.price, 
			t.is_completed,
			u.name, 
			u.surname,
			u.phone_number,
			u.tg_username, 
			u.description,
			r.start_address,
			r.start_latitude,
			r.start_longitude,
			r.stop_address,
			r.stop_latitude,
			r.stop_longitude,
			v.brand,
			v.license_plate,
			(
				SELECT COUNT(*) 
				FROM booking b 
				WHERE b.trip_id = t.id
			) AS bookings_count,
			t.seats_count - (
				SELECT COUNT(*) 
				FROM booking b 
				WHERE b.trip_id = t.id
			) AS available_seats
		FROM 
			trip t
		JOIN 
			"user" u ON t.driver_id = u.id
		JOIN 
			route r ON t.route_id = r.id
		JOIN 
			vehicle v ON t.vehicle_id = v.id
		LEFT JOIN 
			booking b ON t.id = b.trip_id AND b.is_approved = true
		WHERE 
			t.driver_id = $1
	`
	err := rep.db.Select(&trips, query, user_id)
	if err != nil {
		return nil, err
	}
	return trips, nil
}

func (rep *TripRepository) UpdateTripComplete(trip_id int) error {
	query := `UPDATE "trip" SET is_completed = true WHERE id = $1`
	_, err := rep.db.Exec(query, trip_id)
	return err
}

func (rep *TripRepository) DeleteTripById(trip_id int) error {
	query_book := `DELETE FROM "booking" WHERE trip_id = $1`
	_, err_b := rep.db.Exec(query_book, trip_id)
	if err_b != nil && err_b != sql.ErrNoRows {
		return err_b
	}
	query := `DELETE FROM "trip" WHERE id = $1`
	_, err := rep.db.Exec(query, trip_id)
	return err
}
