package main

import (
	"fmt"
	"sa_web_service/internal/models"
	"sa_web_service/internal/database"
	"sa_web_service/internal/handlers"
	"sa_web_service/internal/handlers/middlewares"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	env := &models.ENV{}

	env.Load(".env")

	db := getDB(env)

	autoFunc(db, env)

	fmt.Printf("[MAIN] Subiendo servidor en puerto %s\n",env.PORT)

	serverUp(db,env)
}

func serverUp(db *gorm.DB, env *models.ENV){
	r := gin.Default()
	
	r.Use(middlewares.GinContextToContext())

	r.POST("/query", handlers.GraphQL(db,env))
	r.GET("/", handlers.Playground())
	
	r.Run()
}

func getDB (env *models.ENV) (db *gorm.DB){
	logMode := logger.Error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", env.DB_HOST, env.DB_USER, env.DB_PASSWORD, env.DB_NAME, env.DB_PORT, env.DB_SSLMODE)

	if env.LOG_MODE{
		logMode = logger.Info
	}

	db,err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})

	if err != nil {
		panic(err)
	}


	return
}

func autoFunc(db *gorm.DB, env *models.ENV) {

	if env.AUTO_MIGRATE {
		database.ExecAll(db)
	}

	if env.GIN_MODE == "release"{
		gin.SetMode(gin.ReleaseMode)
	}

}

