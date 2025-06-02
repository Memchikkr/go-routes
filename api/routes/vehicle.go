package routes

import (
	"github.com/Memchikkr/go-routes/api/controllers"
	"github.com/Memchikkr/go-routes/bootstrap"
	"github.com/Memchikkr/go-routes/internal/repositories"
	"github.com/Memchikkr/go-routes/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewVehicleRouter(env *bootstrap.Env, db *sqlx.DB, group *gin.RouterGroup) {
	vr := repositories.NewVehicleRepository(db)
	vc := &controllers.VehicleController{
		Service: services.NewVehicleService(vr, env),
		Env: env,
	}
	group.POST("/vehicles", vc.CreateVehicleRecord)
	group.GET("/vehicles", vc.GetVehicleRecords)
	group.DELETE("/vehicles/:id", vc.DeleteVehicleRecord)
}
