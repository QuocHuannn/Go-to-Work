//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/QuocHuannn/Go-to-Work/internal/controller"
	"github.com/QuocHuannn/Go-to-Work/internal/repo"
	"github.com/QuocHuannn/Go-to-Work/internal/service"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// Set provides user dependencies
var userSet = wire.NewSet(
	provideUserRepo,
	provideUserAuthRepo,
	provideUserService,
	provideUserController,
)

func provideUserRepo(db *gorm.DB) repo.IUserRepository {
	return repo.NewUserRepository(db)
}

func provideUserAuthRepo(db *gorm.DB) repo.IUserAuthRepository {
	return repo.NewUserAuthRepository(db)
}

func provideUserService(userRepo repo.IUserRepository, userAuthRepo repo.IUserAuthRepository) service.IUserService {
	return service.NewUserService(userRepo, userAuthRepo)
}

func provideUserController(userService service.IUserService) *controller.UserController {
	return controller.NewUserController(userService)
}

// InitUserRouterHandler initializes user controller with all dependencies
func InitUserRouterHandler(db *gorm.DB) (*controller.UserController, error) {
	wire.Build(userSet)
	return &controller.UserController{}, nil
}
