package database

import (
	"final-project/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "12345"
	dbName   = "final-project"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname= %s sslmode= disable ", host, port, user, password, dbName)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connection to db:", err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.SocialMedia{}, models.Comment{})
}

func GetDB() *gorm.DB {
	return db
}
