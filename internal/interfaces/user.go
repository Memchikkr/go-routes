package interfaces

import (
	"github.com/Memchikkr/go-routes/internal/models"
)

type UserRepository interface {
	GetUserByID(id int) (*models.UserResponse, error)
	PutUserData(id int, data *models.UserPutRequest) (error)
}

type UserService interface {
	GetUserByID(id int) (*models.UserResponse, error)
	PutUserData(id int, data *models.UserPutRequest) (error)
}
