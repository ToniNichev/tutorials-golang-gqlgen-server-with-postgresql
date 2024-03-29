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

type Bookmark struct {
	ID       uint `gorm:"primaryKey"`
	UserId   string
	Name     string
	Group    string
	MetaData pgtype.JSONB `gorm:"type:jsonb" json:"fieldnameofjsonb"`
}

func autoMigrateDB(db *gorm.DB) {
	// Perform database migration
	func() {
		err := db.AutoMigrate(&User{})
		if err != nil {
			fmt.Println(err)
		}
	}()

	func() {
		err := db.AutoMigrate(&Bookmark{})
		if err != nil {
			fmt.Println(err)
		}
	}()
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

// Bookmarks functions

func AddBookmark(user_id string, name string, group string, metaData string) (*Bookmark, error) {
	jsonData := pgtype.JSONB{}
	err := jsonData.Set([]byte(metaData))
	if err != nil {
		return nil, err
	}
	db, err := connectToPostgreSQL()
	if err != nil {
		return nil, err
	}

	// Create a user
	bookmark := Bookmark{UserId: user_id, Name: name, Group: group, MetaData: jsonData}

	result := db.Create(&bookmark)
	if result.Error != nil {
		return nil, result.Error
	}

	return &bookmark, nil
}

func GetBookmarks(user_id string) (*[]Bookmark, error) {
	db, err := connectToPostgreSQL()
	if err != nil {
		return nil, err
	}

	var bookmark []Bookmark
	//result := db.Find(&bookmark)

	result := db.Table("bookmarks").Where("user_id=?", user_id).Find(&bookmark)

	if result.Error != nil {
		return nil, result.Error
	}

	return &bookmark, nil
}

// USER functions

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

func CreateDB() error {
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

func GetUserByMetaData(metaDataFilter string) (*[]User, error) {
	db, err := connectToPostgreSQL()
	if err != nil {
		return nil, err
	}

	var user []User
	// result := db.First(&user, userID)

	result := db.Where(metaDataFilter).Find(&user)
	fmt.Println(user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
