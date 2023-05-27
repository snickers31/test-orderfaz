package db

import (
	"log"

	"github.com/snickers31/test-orderfaz/auth-svc/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open("host=localhost user=rizal31 password=oumendt31 dbname=orderfaz port=5432 sslmode=disable TimeZone=Asia/Jakarta client_encoding=UTF8"))

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{})
	return Handler{db}
}
