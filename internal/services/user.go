package services

import (
	"github.com/Memchikkr/go-routes/bootstrap"
	"github.com/Memchikkr/go-routes/internal/interfaces"
	"github.com/Memchikkr/go-routes/internal/models"
)

type UserService struct {
	repository interfaces.UserRepository
	env        *bootstrap.Env
}

func NewUserService(repository interfaces.UserRepository, env *bootstrap.Env) interfaces.UserService {
	return &UserService{repository: repository, env: env}
}

func (service *UserService) GetUserByID(id int) (*models.UserResponse, error) {
	user, err := service.repository.GetUserByID(id)

	if err != nil {
		return nil, err
	}
	return user, nil
}


func (service *UserService) PutUserData(id int, data *models.UserPutRequest) (error) {
	err := service.repository.PutUserData(id, data)
	return err
}
