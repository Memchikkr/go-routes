package controllers

import (
	"net/http"
	"strconv"

	"github.com/Memchikkr/go-routes/bootstrap"
	"github.com/Memchikkr/go-routes/internal/interfaces"
	"github.com/Memchikkr/go-routes/internal/models"
	"github.com/gin-gonic/gin"
)

type BookingController struct {
	Service interfaces.BookingService
	Env     *bootstrap.Env
}

// CreateBookingRecord godoc
//
//	@Summary		CreateBookingRecord
//	@Description	Забронировать поездку
//	@Tags			booking
//
// @Security BearerAuth
//
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id поездки"
//	@Success		201		{object}	nil
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/trips/{id}/bookings [post]
func (bc *BookingController) CreateBookingRecord(c *gin.Context) {
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

	tripID := c.Param("id")
	tripIdInt, err := strconv.Atoi(tripID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid trip id",
		})
		return
	}

	err_ins := bc.Service.InsertBookingData(tripIdInt, idInt)
	if err_ins != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create booking record",
		})
		return
	}
	c.JSON(http.StatusCreated, "")
}

// GetBookingRecords godoc
//
//	@Summary		GetBookingRecords
//	@Description	Получить все записи о забронированных местах
//	@Tags			booking
//
// @Security BearerAuth
//
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id поездки"
//	@Success		200		{object}	[]models.BookingWithDetails
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/trips/{id}/bookings [get]
func (tc *BookingController) GetBookingRecords(c *gin.Context) {
	tripID := c.Param("id")
	tripIdInt, err := strconv.Atoi(tripID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid trip id",
		})
		return
	}

	bookings, err := tc.Service.GetBookingRecordsByTripId(tripIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to get bookings",
		})
		return
	}
	if bookings == nil {
		c.JSON(http.StatusOK, []int{})
		return
	}
	c.JSON(http.StatusOK, bookings)
}

// GetMyBookingRecords godoc
//
//	@Summary		GetMyBookingRecords
//	@Description	Получить запись о своих бронированиях
//	@Tags			booking
//
// @Security BearerAuth
//
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	[]models.BookingWithDetails
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/users/me/bookings [get]
func (bc *BookingController) GetMyBookingRecords(c *gin.Context) {
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
	bookings, err := bc.Service.GetBookingRecordsByUserId(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to get your booking records",
		})
		return
	}
	if bookings == nil {
		c.JSON(http.StatusOK, []int{})
		return
	}
	c.JSON(http.StatusOK, bookings)
}

// DeleteBookingRecord godoc
//
//	@Summary		DeleteBookingRecord
//	@Description	Удалить запись о бронировании
//	@Tags			booking
//
// @Security BearerAuth
//
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id поездки"
//	@Success		200		{object}	nil
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/trips/{id}/bookings [delete]
func (bc *BookingController) DeleteBookingRecord(c *gin.Context) {
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

	upd_err := bc.Service.DeleteBookingRecord(tripIdInt, idInt)
	if upd_err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to delete booking record",
		})
		return
	}
	c.JSON(http.StatusOK, "")
}
