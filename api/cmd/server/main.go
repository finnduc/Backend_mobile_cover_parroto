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
	"go-cover-parroto/internal/modules/user"

	_ "go-cover-parroto/cmd/server/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Engflix API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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

	r.GET("/swagger/*any", func(c *gin.Context) {
		if c.Param("any") == "/" || c.Param("any") == "" {
			c.Redirect(302, "/swagger/index.html")
			return
		}
		ginSwagger.WrapHandler(swaggerFiles.Handler)(c)
	})

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
		db := database.DB
		auth.RegisterRoutes(v1, db)
		user.RegisterRoutes(v1, db)
		lesson.RegisterRoutes(v1, db)
		category.RegisterRoutes(v1, db)
		bookmark.RegisterRoutes(v1, db)
	}

	log.Printf("API server running, documentation at http://localhost:%s/swagger", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
