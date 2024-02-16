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
}

func connectToPostgreSQL() (*gorm.DB, error) {
	dsn := "user=mynews password=test123 dbname=tests host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetUserByID(userID uint) (*User, error) {
	db, err := connectToPostgreSQL()
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

func GetUserByMetaData(metaDataFilter string) (*User, error) {
	db, err := connectToPostgreSQL()
	if err != nil {
		return nil, err
	}

	var user User
	// result := db.First(&user, userID)

	result := db.Where(metaDataFilter).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
