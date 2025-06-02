package routes

import (
	"github.com/Memchikkr/go-routes/api/controllers"
	"github.com/Memchikkr/go-routes/bootstrap"
	"github.com/Memchikkr/go-routes/internal/repositories"
	"github.com/Memchikkr/go-routes/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewTripRouter(env *bootstrap.Env, db *sqlx.DB, group *gin.RouterGroup) {
	tr := repositories.NewTripRepository(db)
	tc := &controllers.TripController{
		Service: services.NewTripService(tr, env),
		Env:     env,
	}
	group.POST("/trips", tc.CreateTrip)
	group.GET("/trips", tc.GetTrips)
	group.GET("/users/:id/trips", tc.GetMyTrips)
	group.GET("/trips/:id", tc.GetTrip)
	group.POST("/trips/:id/complete", tc.CompleteTrip)
	group.DELETE("/trips/:id", tc.DeleteTrip)
}
