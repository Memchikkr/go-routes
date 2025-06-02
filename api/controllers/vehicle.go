package controllers

import (
	"net/http"
	"strconv"

	"github.com/Memchikkr/go-routes/bootstrap"
	"github.com/Memchikkr/go-routes/internal/interfaces"
	"github.com/Memchikkr/go-routes/internal/models"
	"github.com/gin-gonic/gin"
)

type VehicleController struct {
	Service interfaces.VehicleService
	Env     *bootstrap.Env
}

// CreateVehicleRecord godoc
//
//	@Summary		CreateVehicleRecord
//	@Description	Создать запись об автомобиле
//	@Tags			vehicle
// @Security BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			input	body		models.VehicleCreateRequest	true	"Данные для создания"
//	@Success		201		{object}	nil
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/vehicles [post]
func (vc *VehicleController) CreateVehicleRecord(c *gin.Context) {
	var request models.VehicleCreateRequest
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
	err := vc.Service.InsertVehicleData(idInt, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create vehicle record",
		})
		return
	}
	c.JSON(http.StatusCreated, "")
}

// GetVehicleRecords godoc
//
//	@Summary		GetVehicleRecords
//	@Description	Получить записи об автомобилях
//	@Tags			vehicle
// @Security BearerAuth
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	[]models.Vehicle
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/vehicles [get]
func (vc *VehicleController) GetVehicleRecords(c *gin.Context) {
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
	vehicles, err := vc.Service.GetVehiclesByUserId(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to get vehicle records",
		})
		return
	}
	if vehicles == nil {
		c.JSON(http.StatusOK, []int{})
		return
	}
	c.JSON(http.StatusOK, vehicles)
}

// DeleteVehicleRecord godoc
//
//	@Summary		DeleteVehicleRecord
//	@Description	Удалить запись об автомобиле
//	@Tags			vehicle
// @Security BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id транспорта"
//	@Success		200		{object}	nil
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/vehicle/{id} [delete]
func (vc *VehicleController) DeleteVehicleRecord(c *gin.Context) {
	vehicleID := c.Param("id")
	vehicleIdInt, err := strconv.Atoi(vehicleID)
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

	upd_err := vc.Service.DeleteVehicleById(vehicleIdInt, idInt)
	if upd_err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to delete vehicle record",
		})
		return
	}
	c.JSON(http.StatusOK, "")
}
