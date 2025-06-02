package routes

import (
	"github.com/Memchikkr/go-routes/api/controllers"
	"github.com/Memchikkr/go-routes/bootstrap"
	"github.com/Memchikkr/go-routes/internal/repositories"
	"github.com/Memchikkr/go-routes/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewUserRouter(env *bootstrap.Env, db *sqlx.DB, group *gin.RouterGroup) {
	ur := repositories.NewUserRepository(db)
	uc := &controllers.UserController{
		Service: services.NewUserService(ur, env),
		Env: env,
	}
	group.GET("/users/me", uc.GetMe)
	group.PUT("/users/me", uc.UpdateUser)
}
