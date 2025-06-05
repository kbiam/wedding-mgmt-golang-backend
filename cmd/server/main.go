package main

import (
	"gorm/api"
	"gorm/config"
	"gorm/db"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main() {
	db.Connect()
	config.LoadEnv()

	originUrl := os.Getenv("ORIGIN_URL")

	router := gin.Default()
	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{originUrl},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        AllowCredentials: true,
    }))
	api.SetupRoutes(router)
	router.Run("0.0.0.0:" + os.Getenv("PORT"))
}