package main

import (
	"log"
	"time"

	"go-cover-parroto/internal/configs"
	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/modules/auth"
	"go-cover-parroto/internal/modules/category"
	"go-cover-parroto/internal/modules/chat"
	"go-cover-parroto/internal/modules/dictation_status"
	"go-cover-parroto/internal/modules/learning_history"
	"go-cover-parroto/internal/modules/lesson"
	"go-cover-parroto/internal/modules/pronunciation"
	"go-cover-parroto/internal/modules/shadowing_status"
	"go-cover-parroto/internal/modules/transcript"
	transcriptbookmark "go-cover-parroto/internal/modules/transcript_bookmark"
	"go-cover-parroto/internal/modules/transcript_progress"
	"go-cover-parroto/internal/modules/user"
	vocabcat "go-cover-parroto/internal/modules/vocabulary_category"
	vocabdeck "go-cover-parroto/internal/modules/vocabulary_deck"
	vocabitem "go-cover-parroto/internal/modules/vocabulary_item"

	_ "go-cover-parroto/cmd/server/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Engflix API
// @version         1.0
// @description     Engflix - Language Learning Platform API
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3001
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	cfg := configs.Load()

	err := logger.Init(cfg.Logger)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.S().Sync()

	if err := database.Init(cfg.Postgres); err != nil {
		logger.S().Fatalf("Failed to connect to database: %v", err)
	}

	if err != nil {
		logger.S().Fatalf("WARNING: Failed to initialize Firebase (%v) — continuing without auth", err)
	}

	port := cfg.Server.Port

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
		auth.RegisterRoutes(v1)
		user.RegisterRoutes(v1)
		lesson.RegisterRoutes(v1, db)
		category.RegisterRoutes(v1, db)
		transcriptbookmark.RegisterRoutes(v1, db)
		transcript.RegisterRoutes(v1, db)
		shadowing_status.RegisterRoutes(v1, db)
		transcript_progress.RegisterRoutes(v1, db)
		dictation_status.RegisterRoutes(v1, db)
		learning_history.RegisterRoutes(v1, db)
		vocabcat.RegisterRoutes(v1, db)
		vocabdeck.RegisterRoutes(v1, db)
		vocabitem.RegisterRoutes(v1, db)
		chat.RegisterRoutes(v1, db)
		pronunciation.RegisterRoutes(v1, db)
	}

	logger.S().Infof("API server running, documentation at http://localhost:%s/swagger", port)
	if err := r.Run(":" + port); err != nil {
		logger.S().Fatalf("Failed to start server: %v", err)
	}
}
