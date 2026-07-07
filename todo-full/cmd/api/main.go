package main

import (
	"log"
	"todo-full/internal/config"
	"todo-full/internal/database"
	"todo-full/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration", err)
	}

	pool, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to load db", err)
	}

	defer pool.Close()

	var router *gin.Engine = gin.Default()

	router.SetTrustedProxies(nil)
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "todo api is running",
			"status":  "success",
		})
	})
	router.POST("/todos", handlers.CreateTodoHandler(pool))

	router.Run(":" + cfg.Port)
}
