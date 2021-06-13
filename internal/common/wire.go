//go:generate wire
//+build wireinject

package internal

import (
	controller2 "easygo/internal/controller"
	dao2 "easygo/internal/dao"
	model2 "easygo/internal/model"
	router2 "easygo/internal/router"
	"easygo/pkg/jwt"
	"github.com/google/wire"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InitGorm,
		InitRedis,
		controller2.ControllerSet,
		model2.ModelSet,
		dao2.DaoSet,
		router2.RouterSet,
		InitGinEngine,
		InjectorSet,
		jwt.JWTSet,
	)

	return nil, nil, nil
}
