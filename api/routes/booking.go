package routes

import (
	"github.com/Memchikkr/go-routes/api/controllers"
	"github.com/Memchikkr/go-routes/bootstrap"
	"github.com/Memchikkr/go-routes/internal/repositories"
	"github.com/Memchikkr/go-routes/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewBookingRouter(env *bootstrap.Env, db *sqlx.DB, group *gin.RouterGroup) {
	br := repositories.NewBookingRepository(db)
	tr := repositories.NewTripRepository(db)
	bc := &controllers.BookingController{
		Service: services.NewBookingService(br, tr, env),
		Env:     env,
	}
	group.POST("/trips/:id/bookings", bc.CreateBookingRecord)
	group.GET("/trips/:id/bookings", bc.GetBookingRecords)
	group.GET("/users/me/bookings", bc.GetMyBookingRecords)
	group.DELETE("/trips/:id/bookings", bc.DeleteBookingRecord)
}
