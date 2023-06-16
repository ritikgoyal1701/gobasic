package Router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	experiment "gorm/Experiment"
	"gorm/models"
)

func GetDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=books port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&models.User{})
	return db
}

func InitalizeRouter() {

	router := gin.Default()
	userController := experiment.Wire(GetDB())
	// Set up a group of routes
	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", userController.GetUser)
		v1.POST("/user", userController.CreateUser)
		v1.POST("/users", userController.CreateUserAll)
		v1.PUT("/user", userController.UpsertUser)
		v1.PUT("/users", userController.UpdateUser)
		v1.DELETE("/users", userController.DeleteUser)
	}

	router.Run(":8080")

}
