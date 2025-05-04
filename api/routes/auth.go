package routes

import (
	"github.com/Memchikkr/go-routes/api/controllers"
	"github.com/Memchikkr/go-routes/bootstrap"
	"github.com/Memchikkr/go-routes/internal/repositories"
	"github.com/Memchikkr/go-routes/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewAuthRouter(env *bootstrap.Env, db *sqlx.DB, group *gin.RouterGroup) {
	ar := repositories.NewAuthRepository(db)
	ac := &controllers.AuthController{
		Service: services.NewAuthService(ar, env),
		Env: env,
	}
	group.POST("/auth", ac.Auth)
	group.POST("/auth/refresh", ac.Refresh)
}
