//go:build wireinject
package wire

import (
	"go-cover-parroto/internal/controller"
	"go-cover-parroto/internal/repo"
	"go-cover-parroto/internal/service"

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

