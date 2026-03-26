//go:build wireinject
package wire

import (
	"go-familytree/internal/controller"
	"go-familytree/internal/repo"
	"go-familytree/internal/service"

	"github.com/google/wire"
)

func InitUserRouterHandle() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepo,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}

