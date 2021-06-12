package model

import "github.com/google/wire"

// ModelSet model 注入 DI
var ModelSet = wire.NewSet(
	UserModelSet,
)
