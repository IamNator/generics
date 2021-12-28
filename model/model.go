package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type (
	Password        string
	TransactionType string
	User            struct {
		gorm.Model
		FirstName string   `json:"first_name" binding:"required"`
		LastName  string   `json:"last_name" binding:"required"`
		Email     string   `json:"email" binding:"required"`
		Password  Password `json:"password" binding:"required"`
	}

	Transaction struct {
		gorm.Model
		Amount    float32         `json:"amount"`
		Type      TransactionType `json:"type"`
		Reference string          `json:"reference"`
	}
)

func (User) TableName() string {
	return "user"
}

func (tt TransactionType) IsValid() bool {
	switch string(tt) {
	case "credit", "debit":
		return true
	default:
		return false
	}
}

func (p Password) String() string {
	return string(p)
}

func (p Password) Hash() Password {
	//check if password is already hashed
	if len(p.String()) > 55 {
		return p
	}

	password := []byte(p.String())
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err.Error())
	}

	return Password(hashedPassword)
}
