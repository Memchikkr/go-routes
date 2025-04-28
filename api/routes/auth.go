package routes

import (
	"github.com/Memchikkr/go-routes/api/controllers"
	"github.com/Memchikkr/go-routes/bootstrap"
	"github.com/gin-gonic/gin"
)

func NewAuthRouter(env *bootstrap.Env, group *gin.RouterGroup) {
	ac := &controllers.AuthController{
		Env: env,
	}
	group.POST("/auth", ac.Auth)
}
