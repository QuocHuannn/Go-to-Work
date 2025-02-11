// go:build wireinject
package wire

import (
	"github.com/QuocHuannn/Go-to-Work/internal/controller"
	"github.com/QuocHuannn/Go-to-Work/internal/repo"
	"github.com/QuocHuannn/Go-to-Work/internal/service"
	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)
	return new(controller.UserController), nil
}
