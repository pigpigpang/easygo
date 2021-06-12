//go:generate wire
//+build wireinject

package internal

import (
	"easygo/internal/controller"
	"easygo/internal/dao"
	"easygo/internal/model"
	"easygo/internal/router"
	"easygo/pkg/jwt"
	"github.com/google/wire"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InitGorm,
		InitRedis,
		controller.ControllerSet,
		model.ModelSet,
		dao.DaoSet,
		router.RouterSet,
		InitGinEngine,
		InjectorSet,
		jwt.JWTSet,
	)

	return nil, nil, nil
}
