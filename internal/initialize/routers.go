package initialize

import (
	"go-cover-parroto/global"
	"go-cover-parroto/internal/middlewares"
	"go-cover-parroto/internal/routers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(deps *AppDependencies) *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// Global Middlewares
	r.Use(gin.Recovery())
	r.Use(middlewares.RequestIDMiddleware())
	r.Use(middlewares.LoggerMiddleware())
	r.Use(middlewares.CORSMiddleware())

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	MainGroup := r.Group("/api/v1")
	{
		MainGroup.GET("/checkstatus", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})
	}

	routers.RouterGroupApp.InitRouterGroup(MainGroup,
		deps.AuthCtrl,
		deps.LessonCtrl,
		deps.AttemptCtrl,
		deps.AnswerCtrl,
		deps.ProgressCtrl,
		deps.BookmarkCtrl,
		deps.CategoryCtrl,
	)

	return r
}
