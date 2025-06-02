package controllers

import (
	"net/http"

	"github.com/Memchikkr/go-routes/bootstrap"
	"github.com/Memchikkr/go-routes/internal/interfaces"
	"github.com/Memchikkr/go-routes/internal/models"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service interfaces.UserService
	Env     *bootstrap.Env
}

// GetMe godoc
//
//	@Summary		GetMe
//	@Description	Возвращает сущность пользователя
//	@Tags			user
// @Security BearerAuth
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	models.UserResponse
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/users/me [get]
func (uc *UserController) GetMe(c *gin.Context) {

	id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found"})
		return
	}

	idInt, ok := id.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}
	user, err := uc.Service.GetUserByID(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
            Code:    http.StatusInternalServerError,
            Message: "Failed to get user",
        })
        return
	}
	c.JSON(http.StatusOK, user)
}


// UpdateUser godoc
//
//	@Summary		UpdateUser
//	@Description	Обновляет данные пользователя
//	@Tags			user
// @Security BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			input	body		models.UserPutRequest	true	"Данные для обновления"
//	@Success		200		{object}	nil
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/users/me [put]
func (uc *UserController) UpdateUser(c *gin.Context) {

	id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found"})
		return
	}

	idInt, ok := id.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}

	var request models.UserPutRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{
            Code:    http.StatusBadRequest,
            Message: "Invalid request",
        })
        return
    }

	err := uc.Service.PutUserData(idInt, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
            Code:    http.StatusInternalServerError,
            Message: "Failed to update user",
        })
        return
	}
	c.JSON(http.StatusOK, "")
}
