package repositories

import (
	"github.com/jmoiron/sqlx"

	"github.com/Memchikkr/go-routes/internal/interfaces"
	"github.com/Memchikkr/go-routes/internal/models"
)

type BookingRepository struct {
	db *sqlx.DB
}

func NewBookingRepository(db *sqlx.DB) interfaces.BookingRepository {
	return &BookingRepository{db: db}
}

func (rep *BookingRepository) GetBookingRecordByTripIdAndUserId(trip_id int, user_id int) (*models.Booking, error) {
	var booking models.Booking
	query := `SELECT * FROM "booking" WHERE trip_id = $1 and user_id = $2`
	err := rep.db.Get(&booking, query, trip_id, user_id)

	if err != nil {
		return nil, err
	}

	return &booking, nil
}

func (rep *BookingRepository) InsertBookingData(trip_id int, user_id int) (error) {
	query := `INSERT INTO "booking" (trip_id, user_id, booking_time, is_approved)
				VALUES ($1, $2, CURRENT_TIMESTAMP, $3)`
	_, err := rep.db.Exec(query, trip_id, user_id, true)

	return err
}

func (rep *BookingRepository) GetBookingRecordsByTripId(trip_id int) ([]*models.BookingWithDetails, error) {
	var bookings []*models.BookingWithDetails
	query := `
		SELECT 
			b.trip_id, 
			b.user_id,
			b.booking_time,
			b.is_approved,
			t.departure_time, 
			t.arrival_time, 
			t.price, 
			t.is_completed,
			t.driver_id,
			u.tg_username, 
			u.phone_number,
			r.start_address,
			r.start_latitude,
			r.start_longitude,
			r.stop_address,
			r.stop_latitude,
			r.stop_longitude,
			v.brand,
			v.license_plate
		FROM 
			"booking" b
		JOIN 
			trip t ON b.trip_id = t.id AND t.id = $1
		JOIN 
			"user" u ON b.user_id = u.id
		JOIN 
			route r ON t.route_id = r.id
		JOIN 
			vehicle v ON t.vehicle_id = v.id
		WHERE 
			1=1
	`
	err := rep.db.Select(&bookings, query, trip_id)
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func (rep *BookingRepository) GetBookingRecordsByUserId(user_id int) ([]*models.BookingWithDetails, error) {
	var bookings []*models.BookingWithDetails
	query := `
		SELECT 
			b.trip_id, 
			b.user_id,
			b.booking_time,
			b.is_approved,
			t.departure_time, 
			t.arrival_time, 
			t.price, 
			t.is_completed,
			t.driver_id,
			u.tg_username, 
			u.phone_number,
			r.start_address,
			r.start_latitude,
			r.start_longitude,
			r.stop_address,
			r.stop_latitude,
			r.stop_longitude,
			v.brand,
			v.license_plate
		FROM 
			"booking" b
		JOIN 
			trip t ON b.trip_id = t.id
		JOIN 
			"user" u ON t.driver_id = u.id
		JOIN 
			route r ON t.route_id = r.id
		JOIN 
			vehicle v ON t.vehicle_id = v.id
		WHERE 
			b.user_id = $1
	`
	err := rep.db.Select(&bookings, query, user_id)
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func (rep *BookingRepository) DeleteBookingRecord(trip_id int, user_id int) (error) {
	query := `DELETE FROM "booking" WHERE trip_id = $1 and user_id = $2`
	_, err := rep.db.Exec(query, trip_id, user_id)
	return err
}
