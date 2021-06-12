package dao

import "github.com/google/wire"

// DaoSet dao 层注入 DI
var DaoSet = wire.NewSet(
	UserDaoSet,
)
