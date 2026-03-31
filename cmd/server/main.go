package main

import (
	_ "go-familytree/docs"
	"go-familytree/internal/initialize"
)

// @title Dictation Learning System API
// @version 1.0
// @description API for learning English via dictation (video-based).
// @host localhost:8002
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	
	initialize.Run()
}

