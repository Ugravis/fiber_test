package database

import (
	"fmt"
	"log"
	"os"

	"fiber-test/internal/models"

	"github.com/fatih/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

var bgGreen = color.New(color.BgGreen, color.FgBlack).SprintFunc()

func ConnectDB() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DEV_DB_USERNAME"),
		os.Getenv("DEV_DB_PASSWORD"),
		os.Getenv("DEV_DB_HOST"),
		os.Getenv("DEV_DB_PORT"),
		os.Getenv("DEV_DB_NAME"),
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
			Colorful: true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatal(`Echec de la connexion BDD:`, err)
	}

	log.Println(bgGreen(fmt.Sprintf("Connexion à la BDD %s avec succès", os.Getenv("DEV_DB_NAME"))))
	
	DB = db

	MigrateDb()
}

func MigrateDb() {
	DB.AutoMigrate(&models.User{ })
	log.Println(bgGreen(`Migration de la BDD avec succès. BDD opérationnelle`))
}