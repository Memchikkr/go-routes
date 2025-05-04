package routes

import (
	"github.com/Memchikkr/go-routes/bootstrap"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Setup(env *bootstrap.Env, db *sqlx.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")
	NewAuthRouter(env, db, publicRouter)
}
