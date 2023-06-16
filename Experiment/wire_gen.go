// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package experiment

import (
	"gorm.io/gorm"
)

// Injectors from wire.go:

func Wire(db *gorm.DB) *Controller {
	repository := NewRepository(db)
	service := NewService(repository)
	controller := NewController(service)
	return controller
}
