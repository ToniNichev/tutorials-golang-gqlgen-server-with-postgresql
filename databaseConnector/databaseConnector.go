package databaseConnector

import (
	"github.com/jackc/pgtype"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Email    string
	Age      int
	MetaData pgtype.JSONB `gorm:"type:jsonb" json:"fieldnameofjsonb"`
	// MetaData JSONB `sql:"type:jsonb"`
}

func GetUserByID(userID uint) (*User, error) {
	dsn := "user=mynews password=test123 dbname=tests host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	var user User
	result := db.First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
