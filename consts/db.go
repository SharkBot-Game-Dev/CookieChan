package consts

import (
	"log"

	"github.com/SharkBot-Game-Dev/CookieChan/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
	log.Print("DBに接続しました。")

	initDB()
}

func initDB() {
	if err := DB.AutoMigrate(
		&models.CookieGameUser{},
		&models.CookieGameAchievement{},
		&models.CookieGameItem{},
	); err != nil {
		log.Fatal(err)
	}
}
