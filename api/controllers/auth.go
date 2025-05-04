package controllers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/Memchikkr/go-routes/bootstrap"
	"github.com/Memchikkr/go-routes/internal/interfaces"
	"github.com/Memchikkr/go-routes/internal/models"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
    Service interfaces.AuthService
	Env *bootstrap.Env
}

// Auth godoc
//	@Summary		Authentification
//	@Description	Выполняет аутентификацию и возвращает токен
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			input	body		models.AuthRequest	true	"Данные для входа"
//	@Success		200		{object}	models.AuthResponse
//	@Failure		400		{object}	models.ErrorResponse
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/auth [post]
func (ac *AuthController) Auth(c *gin.Context) {
	var request models.AuthRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{
            Code:    http.StatusBadRequest,
            Message: "Invalid request",
        })
        return
    }
    if request.Id <= 0 {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{
            Code:    http.StatusBadRequest,
            Message: "Telegram ID is required",
        })
        return
    }
    
    user, err := ac.Service.GetUserByTelegramID(request.Id)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            user, err = ac.Service.CreateUser(&request)
            if err != nil {
                c.JSON(http.StatusInternalServerError, models.ErrorResponse{
                    Code:    http.StatusInternalServerError,
                    Message: "Failed to create user",
                })
                return
            }
        } else {
            c.JSON(http.StatusInternalServerError, models.ErrorResponse{
                Code:    http.StatusInternalServerError,
                Message: "Database error",
            })
            return
        }
    }

    tokens, err := ac.Service.GenerateTokens(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{
            Code:    http.StatusInternalServerError,
            Message: "Failed to generate tokens",
        })
        return
    }
    c.JSON(http.StatusOK, tokens)
}

// Refresh godoc
// @Summary Refresh tokens
// @Description	Выполняет аутентификацию с помощью RefreshToken и возвращает новую пару токенов
// @Tags auth
// @Accept json
// @Produce json
// @Param input body models.RefreshRequest true "Refresh token"
// @Success 200 {object} models.AuthResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /auth/refresh [post]
func (ac *AuthController) Refresh(c *gin.Context) {
    var request models.RefreshRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{
            Code:    http.StatusBadRequest,
            Message: "Invalid request",
        })
        return
    }

    tokens, err := ac.Service.RefreshTokens(request.RefreshToken)
    if err != nil {
        c.JSON(http.StatusUnauthorized, models.ErrorResponse{
            Code: http.StatusUnauthorized,
            Message: "invalid refresh token",
        })
        return
    }

    c.JSON(http.StatusOK, tokens)
}
