package databaseConnector

import (
	"fmt"

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

func autoMigrateDB(db *gorm.DB) {
	// Perform database migration
	err := db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println(err)
	}
}

func connectToPostgreSQL() (*gorm.DB, error) {
	// dsn := "user=mynews password=test123 dbname=tests host=localhost port=5432 sslmode=disable"
	dsn := "user=toninichev dbname=tests host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createuserWithMetaData(db *gorm.DB, username string, email string, age int, metaData string) (*User, error) {
	jsonData := pgtype.JSONB{}
	err := jsonData.Set([]byte(metaData))
	if err != nil {
		return nil, err
	}
	// Create a user
	newUser := User{Username: username, Email: email, Age: age, MetaData: jsonData}
	err = createUser(db, &newUser)
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}
func createUser(db *gorm.DB, user *User) error {
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateDB(tableName string) error {
	db, err := connectToPostgreSQL()
	if err != nil {
		return err
	}
	autoMigrateDB(db)
	return nil
}

func CreateUser(username string, email string, age int, metaData string) (*User, error) {
	db, err := connectToPostgreSQL()
	if err != nil {
		return nil, err
	}
	user, err := createuserWithMetaData(db, username, email, age, metaData)
	return user, err
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
