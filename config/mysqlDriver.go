package config

import (
	_mAttendees "altaproject3/features/attendees/data"
	_mComments "altaproject3/features/comments/data"
	_mEvents "altaproject3/features/events/data"
	_mUsers "altaproject3/features/users/data"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dbUsername := os.Getenv("DB_Username")
	dbPassword := os.Getenv("DB_Password")
	dbPort := os.Getenv("DB_Port")
	dbHost := os.Getenv("DB_Host")
	dbName := os.Getenv("DB_Name")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	InitMigrate(db)
	return db
}

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&_mUsers.User{})
	db.AutoMigrate(&_mEvents.Event{})
	db.AutoMigrate(&_mComments.Comment{})
	db.AutoMigrate(&_mAttendees.Attendee{})
}
