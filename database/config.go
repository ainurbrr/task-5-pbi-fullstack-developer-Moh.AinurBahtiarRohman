package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"task-5-vix-btpns-Moh.AinurBahtiarRohman/models"
)

var DB *gorm.DB

func Init() *gorm.DB {
	InitDB()
	InitMigration()
	return DB
}

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func InitDB() {
	config := Config{
		DB_USERNAME: "root",
		DB_PASSWORD: "",
		DB_PORT:     "3306",
		DB_HOST:     "127.0.0.1",
		DB_NAME:     "fpVIXbtpns",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitMigration() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Photo{})

}
