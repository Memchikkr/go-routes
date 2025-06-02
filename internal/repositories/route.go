package repositories

import (
	"github.com/jmoiron/sqlx"

	"github.com/Memchikkr/go-routes/internal/interfaces"
	"github.com/Memchikkr/go-routes/internal/models"
)

type RouteRepository struct {
	db *sqlx.DB
}

func NewRouteRepository(db *sqlx.DB) interfaces.RouteRepository {
	return &RouteRepository{db: db}
}

func (rep *RouteRepository) GetRouteById(route_id int, user_id int) (*models.Route, error) {
	var route models.Route
	query := `SELECT * FROM "route" WHERE id = $1 and user_id = $2`
	err := rep.db.Get(&route, query, route_id, user_id)

	if err != nil {
		return nil, err
	}

	return &route, nil
}

func (rep *RouteRepository) InsertRouteData(user_id int, data *models.RouteCreateRequest) (error) {
	query := `INSERT INTO "route" (user_id, start_address, start_latitude, start_longitude, stop_address, stop_latitude, stop_longitude)
				VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := rep.db.Exec(query, user_id, data.StartAddress, data.StartLatitude, data.StartLongitude, data.StopAddress, data.StopLatitude, data.StopLongitude)

	return err
}

func (rep *RouteRepository) GetRoutesByUserId(user_id int) ([]*models.Route, error) {
	var routes []*models.Route
	query := `SELECT * FROM "route" WHERE user_id = $1`
	err := rep.db.Select(&routes, query, user_id)
	if err != nil {
		return nil, err
	}
	return routes, nil
}

func (rep *RouteRepository) DeleteRouteById(route_id int) (error) {
	query := `DELETE FROM "route" WHERE id = $1`
	_, err := rep.db.Exec(query, route_id)
	return err
}