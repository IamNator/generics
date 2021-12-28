package storage

import (
	"generics/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Storage struct {
	Db *gorm.DB
}

func New() *Storage {
	db, err := gorm.
		Open("postgres", "postgres://ajotxkyerxugmt:002367e4269f2862de7553ce0e1161e68250678d04172935cee2d04fc2f9e34e@ec2-52-213-119-221.eu-west-1.compute.amazonaws.com:5432/d9i56c8ddu2l4n")
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
