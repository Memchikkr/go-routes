package controllers

import (
	"net/http"
	"strconv"

	"github.com/Memchikkr/go-routes/bootstrap"
	"github.com/Memchikkr/go-routes/internal/interfaces"
	"github.com/Memchikkr/go-routes/internal/models"
	"github.com/gin-gonic/gin"
)

type RouteController struct {
	Service interfaces.RouteService
	Env     *bootstrap.Env
}

// CreateRouteRecord godoc
//
//	@Summary		CreateRouteRecord
//	@Description	Создать запись о маршруте
//	@Tags			routes
//
// @Security BearerAuth
//
//	@Accept			json
//	@Produce		json
//	@Param			input	body		models.RouteCreateRequest	true	"Данные для создания"
//	@Success		201		{object}	nil
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/routes [post]
func (rc *RouteController) CreateRouteRecord(c *gin.Context) {
	var request models.RouteCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request",
		})
		return
	}

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
	err := rc.Service.InsertRouteData(idInt, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create route record",
		})
		return
	}
	c.JSON(http.StatusCreated, "")
}

// GetRouteRecords godoc
//
//	@Summary		GetRouteRecords
//	@Description	Получить записи о маршруте
//	@Tags			routes
//
// @Security BearerAuth
//
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	[]models.Route
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/routes [get]
func (rc *RouteController) GetRouteRecords(c *gin.Context) {
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
	routes, err := rc.Service.GetRoutesByUserId(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to get routes records",
		})
		return
	}
	if routes == nil {
		c.JSON(http.StatusOK, []int{})
		return
	}
	c.JSON(http.StatusOK, routes)
}

// DeleteRouteRecord godoc
//
//	@Summary		DeleteRouteRecord
//	@Description	Удалить запись о маршруте
//	@Tags			routes
//
// @Security BearerAuth
//
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id маршрута"
//	@Success		200		{object}	nil
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/routes/{id} [delete]
func (rc *RouteController) DeleteRouteRecord(c *gin.Context) {
	routeID := c.Param("id")
	routeIdInt, err := strconv.Atoi(routeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid trip id",
		})
		return
	}

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

	del_err := rc.Service.DeleteRouteById(routeIdInt, idInt)
	if del_err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to delete route record",
		})
		return
	}
	c.JSON(http.StatusOK, "")
}
