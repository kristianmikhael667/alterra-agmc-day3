package config

import (
	"fmt"
	"main/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	dbusername := os.Getenv("DB_USERNAME")
	dbpass := os.Getenv("DB_PASSWORD")
	dbport := os.Getenv("DB_PORT")
	dbhost := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")

	config := map[string]string{
		"DB_USERNAME": dbusername,
		"DB_PASSWORD": dbpass,
		"DB_PORT":     dbport,
		"DB_HOST":     dbhost,
		"DB_NAME":     dbname,
	}
	var err error
	connectionString := fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config["DB_USERNAME"], config["DB_HOST"], config["DB_PORT"], config["DB_NAME"])

	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
}
