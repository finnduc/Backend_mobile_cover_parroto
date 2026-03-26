package initialize

import (
	"fmt"
	"go-familytree/global"

	"go.uber.org/zap"
)

func Run() {
	// 1. Load configuration
	LoadConfig()

	// 2. Init components
	InitPostgres()
	InitRedis()

	// 3. Wire dependencies
	deps := WireDependencies()

	// 4. Setup router
	r := InitRouter(deps)

	// 5. Start server
	port := global.Config.Server.Port
	global.Logger.Info(fmt.Sprintf("Server is running on port %d", port))
	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		global.Logger.Fatal("Server failed to start", zap.Error(err))
	}
}