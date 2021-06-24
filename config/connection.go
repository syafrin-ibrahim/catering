package config

import (
	"catering/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error

func Init() {

	DB, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/catering?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Println("Connection Failed", err)

	} else {
		log.Println("Connection Success")
	}

	DB.AutoMigrate(&models.User{}, &models.Regency{}, &models.Transaction{}, &models.Paket{})
	DB.Model(&models.Transaction{}).AddForeignKey("regency_id", "regencies(id)", "CASCADE", "CASCADE")
	DB.Model(&models.Transaction{}).AddForeignKey("paket_id", "pakets(id)", "CASCADE", "CASCADE")
	DB.Model(&models.Transaction{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

}
