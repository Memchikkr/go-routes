package controllers

import (
	"net/http"
	"strconv"

	"github.com/Memchikkr/go-routes/bootstrap"
	"github.com/Memchikkr/go-routes/internal/interfaces"
	"github.com/Memchikkr/go-routes/internal/models"
	"github.com/gin-gonic/gin"
)

type TripController struct {
	Service interfaces.TripService
	Env     *bootstrap.Env
}

// CreateTrip godoc
//
//	@Summary		CreateTrip
//	@Description	Создать запись о поездке
//	@Tags			trip
//
// @Security BearerAuth
//
//	@Accept			json
//	@Produce		json
//	@Param			input	body		models.TripCreateRequest	true	"Данные для создания"
//	@Success		201		{object}	nil
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/trips [post]
func (tc *TripController) CreateTrip(c *gin.Context) {
	var request models.TripCreateRequest
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
	err := tc.Service.InsertTripData(idInt, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create trip",
		})
		return
	}
	c.JSON(http.StatusCreated, "")
}

// GetTrips godoc
//
//	@Summary		GetTrips
//	@Description	Получить все записи о поездках
//	@Tags			trip
//
// @Security BearerAuth
//
//	@Accept			json
//	@Produce		json
//
// @Param from      query string  false "Город отправления"  example(Москва)
// @Param to        query string  false "Город назначения"   example(Санкт-Петербург)
// @Param date      query string  false "Дата (YYYY-MM-DD)"  example(2025-05-20)
// @Param min_seats query int     false "Минимальное кол-во мест" minimum(1) example(2)
// @Param max_price query number  false "Максимальная цена"   minimum(0) example(1500.50)
//
//	@Success		200		{object}	[]models.TripWithDetails
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/trips [get]
func (tc *TripController) GetTrips(c *gin.Context) {
	minSeats, err := strconv.Atoi(c.DefaultQuery("min_seats", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid min_seats value"})
		return
	}

	maxPrice, err := strconv.ParseFloat(c.Query("max_price"), 64)
	if err != nil && c.Query("max_price") != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid max_price value"})
		return
	}
	filters := models.TripFilters{
		From:     c.Query("from"),
		To:       c.Query("to"),
		Date:     c.Query("date"),
		MinSeats: minSeats,
		MaxPrice: maxPrice,
	}
	trips, err := tc.Service.GetAllTrips(&filters)
	if err != nil {
		println(err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to get trips",
		})
		return
	}
	if trips == nil {
		c.JSON(http.StatusOK, []int{})
		return
	}
	c.JSON(http.StatusOK, trips)
}

// GetMyTrips godoc
//
//	@Summary		GetMyTrips
//	@Description	Получить все свои записи о поездках
//	@Tags			trip
//
// @Security BearerAuth
//
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	[]models.TripWithDetails
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/users/{id}/trips [get]
func (tc *TripController) GetMyTrips(c *gin.Context) {
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
	trips, err := tc.Service.GetAllMyTrips(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to get your trips",
		})
		return
	}
	if trips == nil {
		c.JSON(http.StatusOK, []int{})
		return
	}
	c.JSON(http.StatusOK, trips)
}

// GetTrip godoc
//
//	@Summary		GetTrip
//	@Description	Получить одну запись о поездке
//	@Tags			trip
//
// @Security BearerAuth
//
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id поездки"
//	@Success		200		{object}	models.Trip
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/trips/{id} [get]
func (tc *TripController) GetTrip(c *gin.Context) {
	tripID := c.Param("id")
	tripIdInt, err := strconv.Atoi(tripID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid trip id",
		})
		return
	}
	trips, err := tc.Service.GetTripById(tripIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to get trip",
		})
		return
	}
	c.JSON(http.StatusOK, trips)
}

// CompleteTrip godoc
//
//	@Summary		CompleteTrip
//	@Description	Завершить поездку
//	@Tags			trip
//
// @Security BearerAuth
//
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id поездки"
//	@Success		200		{object}	nil
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/trips/{id}/complete [post]
func (tc *TripController) CompleteTrip(c *gin.Context) {
	tripID := c.Param("id")
	tripIdInt, err := strconv.Atoi(tripID)
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

	upd_err := tc.Service.UpdateTripComplete(tripIdInt, idInt)
	if upd_err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to complete trip",
		})
		return
	}
	c.JSON(http.StatusOK, "")
}

// DeleteTrip godoc
//
//	@Summary		DeleteTrip
//	@Description	Удалить поездку
//	@Tags			trip
//
// @Security BearerAuth
//
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id поездки"
//	@Success		200		{object}	nil
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/trips/{id} [delete]
func (tc *TripController) DeleteTrip(c *gin.Context) {
	tripID := c.Param("id")
	tripIdInt, err := strconv.Atoi(tripID)
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

	del_err := tc.Service.DeleteTripById(tripIdInt, idInt)
	if del_err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to delete trip",
		})
		return
	}
	c.JSON(http.StatusOK, "")
}
