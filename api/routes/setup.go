package routes

import (
	"github.com/Memchikkr/go-routes/api/middlewares"
	"github.com/Memchikkr/go-routes/bootstrap"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Setup(env *bootstrap.Env, db *sqlx.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")
	NewAuthRouter(env, db, publicRouter)
	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middlewares.JwtAuthMiddleware(env.AccessTokenSecret))
	NewUserRouter(env, db, protectedRouter)
	NewTripRouter(env, db, protectedRouter)
	NewBookingRouter(env, db, protectedRouter)
	NewVehicleRouter(env, db, protectedRouter)
	NewRouteRouter(env, db, protectedRouter)
}
