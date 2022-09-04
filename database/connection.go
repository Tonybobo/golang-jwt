package database

import (
	"github.com/tony/jwt-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	connection, err := gorm.Open(mysql.Open("root:root@(localhost)/jwt-auth?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to jwt-auth")
	}

	DB = connection
	connection.AutoMigrate(&models.User{})
}
