package routes

import (
	"github.com/Memchikkr/go-routes/bootstrap"
	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, gin *gin.Engine) {
	publicRouter := gin.Group("")
	NewAuthRouter(env, publicRouter)
}
