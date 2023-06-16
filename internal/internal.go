package internal

import (
	"github.com/gin-gonic/gin"
	"gorm/models"
)

type UserService interface {
	CreateUser(name, age, phone string) string
	CreateUserAll(name, age, phone []string) string
	GetUser(name, age, phone string) []*models.User
	UpdateUser(nameOrg, ageOrg, phoneOrg, name, age, phone string) string
	DeleteUser(name, age, phone string) string
	UpsertUser(name, age, phone string) string
}

type UserRepository interface {
	CreateUser(user *models.User) bool
	CreateUserAll(users []*models.User) bool
	GetUser(mp map[string]interface{}) []*models.User
	UpdateUser(org map[string]interface{}, update map[string]interface{}) bool
	DeleteUser(mp map[string]interface{}) bool
	UpsertUser(user *models.User) bool
}

type UserController interface {
	CreateUser(ctx *gin.Context)
	CreateUserAll(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	UpdateUser(nameOrg, ageOrg, phoneOrg, name, age, phone string)
	DeleteUser(name, age, phone string)
	UpsertUser(name, age, phone string)
}
