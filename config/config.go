package config

import (
	"user-crud/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitialMigration() {
	DB.AutoMigrate(&models.User{})
}

func InitDB() {
	connectionString := "root:root@tcp(127.0.0.1:3306)/user-crud?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}