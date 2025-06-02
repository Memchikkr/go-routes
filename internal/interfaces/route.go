package interfaces

import "github.com/Memchikkr/go-routes/internal/models"

type RouteRepository interface {
	InsertRouteData(user_id int, data *models.RouteCreateRequest) (error)
	GetRoutesByUserId(user_id int) ([]*models.Route, error)
	GetRouteById(route_id int, user_id int) (*models.Route, error)
	DeleteRouteById(route_id int) (error)
}

type RouteService interface {
	InsertRouteData(user_id int, data *models.RouteCreateRequest) (error)
	GetRoutesByUserId(user_id int) ([]*models.Route, error)
	DeleteRouteById(route_id int, user_id int) (error)
}
