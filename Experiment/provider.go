package experiment

import (
	"github.com/google/wire"
	"gorm/internal"
)

var ProviderSet = wire.NewSet(
	NewController,
	NewService,
	NewRepository,
	wire.Bind(new(internal.UserController), new(*Controller)),
	wire.Bind(new(internal.UserService), new(*Service)),
	wire.Bind(new(internal.UserRepository), new(*Repository)),
)
