package data

import (
	attendees "altaproject3/features/attendees"
	"fmt"

	"gorm.io/gorm"
)

type mysqlAttendeeRepository struct {
	db *gorm.DB
}

func NewAttendeeRepository(conn *gorm.DB) attendees.Data {
	return &mysqlAttendeeRepository{
		db: conn,
	}
}

func (repo *mysqlAttendeeRepository) PostAttendee(data attendees.Core) (int, error) {
	attendee := fromCore(data)
	result := repo.db.Create(&attendee)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to insert data")
	}
	return int(result.RowsAffected), nil
}
