package interfaces

import "github.com/Memchikkr/go-routes/internal/models"

type AuthRepository interface {
	GetUserByTelegramID(tg_id int) (*models.User, error)
	CreateUser(user *models.User) error
}

type AuthService interface {
	GetUserByTelegramID(tg_id int) (*models.User, error)
	CreateUser(auth *models.AuthRequest) (*models.User, error)
	GenerateTokens(user *models.User) (*models.AuthResponse, error)
	RefreshTokens(refreshToken string) (*models.AuthResponse, error)
	VerifyAuthData(auth *models.AuthRequest) (bool, *models.ErrorResponse)
}
