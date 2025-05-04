package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/Memchikkr/go-routes/bootstrap"
	"github.com/Memchikkr/go-routes/internal/interfaces"
	"github.com/Memchikkr/go-routes/internal/models"
)

type AuthService struct {
	repository interfaces.AuthRepository
	env        *bootstrap.Env
}

func NewAuthService(repository interfaces.AuthRepository, env *bootstrap.Env) interfaces.AuthService {
	return &AuthService{repository: repository, env: env}
}

func (service *AuthService) GetUserByTelegramID(tg_id int) (*models.User, error) {
	user, err := service.repository.GetUserByTelegramID(tg_id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *AuthService) CreateUser(auth *models.AuthRequest) (*models.User, error) {
	err := service.repository.CreateUser(&models.User{
		TgUsername: auth.UserName, 
		TgID: auth.Id, 
		Name: sql.NullString{String: auth.FirstName, Valid: true}, 
		SurName: sql.NullString{String: auth.LastName, Valid: true},
	})

	if err != nil {
		return nil, err
	}

	user, err := service.repository.GetUserByTelegramID(auth.Id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *AuthService) GenerateTokens(user *models.User) (*models.AuthResponse, error) {
	return service.GenerateTokenPair(user.ID)
}

func (service *AuthService) RefreshTokens(RefreshToken string) (*models.AuthResponse, error) {
    claims, err := service.ParseToken(RefreshToken, service.env.RefreshTokenSecret)
    if err != nil {
        return nil, err
    }
    userID := claims.(jwt.MapClaims)["sub"].(float64)
    return service.GenerateTokenPair(int(userID))
}

func (service *AuthService) VerifyAuthData(auth *models.AuthRequest) (bool, *models.ErrorResponse) {
	if !isAuthDateValid(auth.AuthDate) {
		return false, nil
	}

	BotToken := service.env.BotToken
	if !verifyTelegramHash(auth, BotToken) {
		return false, nil
	}

	return true, nil
}

func isAuthDateValid(authDate int) bool {
	now := int(time.Now().Unix())
	return now-authDate < 86400
}

func verifyTelegramHash(auth *models.AuthRequest, Token string) bool {
	var Fields []string
	Fields = append(Fields, fmt.Sprintf("auth_date=%d", auth.AuthDate))
	Fields = append(Fields, fmt.Sprintf("first_name=%s", auth.FirstName))
	Fields = append(Fields, fmt.Sprintf("id=%d", auth.Id))
	if auth.LastName != "" {
		Fields = append(Fields, fmt.Sprintf("last_name=%s", auth.LastName))
	}
	if auth.PhotoUrl != "" {
		Fields = append(Fields, fmt.Sprintf("photo_url=%s", auth.PhotoUrl))
	}
	if auth.UserName != "" {
		Fields = append(Fields, fmt.Sprintf("username=%s", auth.UserName))
	}
	sort.Strings(Fields)
	DataCheckString := strings.Join(Fields, "\n")
	SecretKey := sha256.Sum256([]byte(Token))
	HmacHash := hmac.New(sha256.New, SecretKey[:])
	HmacHash.Write([]byte(DataCheckString))
	ExpectedHash := hex.EncodeToString(HmacHash.Sum(nil))
	return ExpectedHash == auth.Hash
}

func (service *AuthService) GenerateTokenPair(userID int) (*models.AuthResponse, error) {
	accessToken, err := generateJWT(
		userID,
		service.env.AccessTokenSecret,
		time.Duration(service.env.AccessTokenExpiryHour)*time.Hour,
	)
	if err != nil {
		return nil, err
	}

	refreshToken, err := generateJWT(
		userID,
		service.env.RefreshTokenSecret,
		time.Duration(service.env.RefreshTokenExpiryHour)*time.Hour,
	)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func generateJWT(userID int, secret string, expiry time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(expiry).Unix(),
		"iat": time.Now().Unix(),
	})

	return token.SignedString([]byte(secret))
}

func (s *AuthService) ParseToken(tokenString, secret string) (jwt.Claims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(secret), nil
    })

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, err
    }

    return token.Claims, nil
}
