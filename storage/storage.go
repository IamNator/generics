package storage

import (
	"generics/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

type Storage struct {
	Db *gorm.DB
}

func New() *Storage {
	db, err := gorm.
		Open("postgres", os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatalln(err.Error())
	}

	return &Storage{
		Db: db,
	}

}

func (s *Storage) InitModels() {
	result := s.Db.AutoMigrate(model.User{})
	if er := result.Error; er != nil {
		log.Fatalln(er.Error())
	}

	result = s.Db.AutoMigrate(model.Transaction{})
	if er := result.Error; er != nil {
		log.Fatalln(er.Error())
	}
}
