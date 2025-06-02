package services

import (
	"github.com/Memchikkr/go-routes/bootstrap"
	"github.com/Memchikkr/go-routes/internal/interfaces"
	"github.com/Memchikkr/go-routes/internal/models"
)

type RouteService struct {
	repository interfaces.RouteRepository
	env        *bootstrap.Env
}

func NewRouteService(repository interfaces.RouteRepository, env *bootstrap.Env) interfaces.RouteService {
	return &RouteService{repository: repository, env: env}
}

func (service *RouteService) InsertRouteData(user_id int, data *models.RouteCreateRequest) (error) {
	err := service.repository.InsertRouteData(user_id, data)
	return err
}

func (service *RouteService) GetRoutesByUserId(user_id int) ([]*models.Route, error) {
	routes, err := service.repository.GetRoutesByUserId(user_id)

	if err != nil {
		return nil, err
	}

	return routes, nil
}

func (service *RouteService) DeleteRouteById(route_id int, user_id int) (error) {
	_, err := service.repository.GetRouteById(route_id, user_id)
	if err != nil {
		return err
	}
	del_err := service.repository.DeleteRouteById(route_id)
	return del_err
}
