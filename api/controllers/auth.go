package controllers

import (
	"net/http"

	"github.com/Memchikkr/go-routes/bootstrap"
	"github.com/Memchikkr/go-routes/internal/models"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Env *bootstrap.Env
}

// Auth godoc
//	@Summary		Аутентификация пользователя
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
	var input models.AuthRequest
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{
            Code:    http.StatusBadRequest,
            Message: "Invalid request",
        })
        return
    }
    
    // token, err := ac.authService.Authenticate(input.Login, input.Password)
    // if err != nil {
    //     c.JSON(http.StatusUnauthorized, models.ErrorResponse{
    //         Code:    http.StatusUnauthorized,
    //         Message: "",
    //     })
    //     return
    // }
	token := "Token"
    c.JSON(http.StatusOK, models.AuthResponse{
        Token: token,
    })
}
