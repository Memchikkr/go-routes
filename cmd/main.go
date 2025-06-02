package main

import (
	"log"

	"github.com/Memchikkr/go-routes/api/routes"
	"github.com/Memchikkr/go-routes/bootstrap"
	"github.com/Memchikkr/go-routes/docs"
	"github.com/Memchikkr/go-routes/migrations"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			API
//	@version		1.0
//	@description	API
//	@host			localhost:3000
//
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Введите: "Bearer {ваш_JWT_токен}"
//
//	@BasePath		/
func main() {
	app := bootstrap.App()
	db := app.DB.DB
	defer app.Close()

	env := app.Env

	migrator := migrations.NewMigrator(
		env.DBHost,
		env.DBPort,
		env.DBUser,
		env.DBPassword,
		env.DBName,
	)

	if err := migrator.Up(); err != nil {
		log.Fatalf("migration error: %v", err)
	}

	gin := gin.Default()

	gin.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
        AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
        ExposeHeaders:   []string{"Content-Length"},
        AllowCredentials: true,
    }))

	docs.SwaggerInfo.BasePath = "/"
	gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.Setup(env, db, gin)

	gin.Run(":" + env.AppPort)
}
