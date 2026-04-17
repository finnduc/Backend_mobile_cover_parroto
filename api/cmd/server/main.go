package main

import (
	"log"
	"time"

	"go-cover-parroto/internal/configs"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/modules/auth"
	"go-cover-parroto/internal/modules/bookmark"
	"go-cover-parroto/internal/modules/category"
	"go-cover-parroto/internal/modules/lesson"
	"go-cover-parroto/internal/modules/progress"
	"go-cover-parroto/internal/modules/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := configs.Load()

	if err := database.Init(cfg.Postgres); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	port := cfg.Server.Port
	if port == "" {
		port = "3001"
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "apikey", "Accept"},
		ExposeHeaders:    []string{"Content-Type", "Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok", "service": "parroto-api"})
		})

		v1 := api.Group("/v1")
		auth.RegisterRoutes(v1)
		user.RegisterRoutes(v1)
		lesson.RegisterRoutes(v1)
		category.RegisterRoutes(v1)
		bookmark.RegisterRoutes(v1)
		progress.RegisterRoutes(v1)
	}

	log.Printf("API server running at http://localhost:%s/api", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
