//go:build wireinject
// +build wireinject

package experiment

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

func Wire(db *gorm.DB) *Controller {
	panic(wire.Build(ProviderSet))
}
