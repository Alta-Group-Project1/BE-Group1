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

func (repo *mysqlAttendeeRepository) DeleteAttendee(id int) (int, error) {
	var data1 Attendee
	tx := repo.db.Where("id = ?", id).Delete(&data1)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to delete data")
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlAttendeeRepository) GetAttendeeByIdEvent(idEvent int) ([]attendees.Core, error) {
	var dataAttend []Attendee
	tx := repo.db.Model(&Attendee{}).Preload("User").Preload("Event").Where("event_id = ?", idEvent).Find(&dataAttend)
	if tx.Error != nil {
		return []attendees.Core{}, tx.Error
	}
	return toCoreList(dataAttend), nil
}

func (repo *mysqlAttendeeRepository) GetAttendeeByIdUser(idUser int) ([]attendees.Core, error) {
	var dataAttend []Attendee
	result := repo.db.Model(&Attendee{}).Preload("User").Preload("Event").Where("user_id = ?", idUser).Find(&dataAttend)
	if result.Error != nil {
		return []attendees.Core{}, result.Error
	}
	return toCoreList(dataAttend), nil
}
