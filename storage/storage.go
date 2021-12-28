package storage

import (
	"generics/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

type Storage struct {
	Db *gorm.DB
}

func New() *Storage {
	db, err := gorm.
		Open("postgres", os.Getenv("DB_DSN"))
	if err != nil {
		panic(err.Error())
	}

	return &Storage{
		Db: db,
	}

}

func (s *Storage) InitModels() {
	result := s.Db.AutoMigrate(model.User{})
	if er := result.Error; er != nil {
		panic(er.Error())
	}

	result = s.Db.AutoMigrate(model.Transaction{})
	if er := result.Error; er != nil {
		panic(er.Error())
	}
}
