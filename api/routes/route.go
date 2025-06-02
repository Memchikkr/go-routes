package routes

import (
	"github.com/Memchikkr/go-routes/api/controllers"
	"github.com/Memchikkr/go-routes/bootstrap"
	"github.com/Memchikkr/go-routes/internal/repositories"
	"github.com/Memchikkr/go-routes/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewRouteRouter(env *bootstrap.Env, db *sqlx.DB, group *gin.RouterGroup) {
	tr := repositories.NewRouteRepository(db)
	tc := &controllers.RouteController{
		Service: services.NewRouteService(tr, env),
		Env:     env,
	}
	group.POST("/routes", tc.CreateRouteRecord)
	group.GET("/routes", tc.GetRouteRecords)
	group.DELETE("/routes/:id", tc.DeleteRouteRecord)
}
