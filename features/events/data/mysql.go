package data

import (
	"altaproject3/features/events"

	"gorm.io/gorm"
)

type mysqlEventRepository struct {
	db *gorm.DB
}

func NewEventRepository(conn *gorm.DB) events.Data {
	return &mysqlEventRepository{
		db: conn,
	}
}

func (repo *mysqlEventRepository) SelectAllEvent(limit, offset uint) (response []events.Core, err error) {
	var dataEvents []Event
	result := repo.db.Preload("User").Find(&dataEvents)
	if result.Error != nil {
		return nil, result.Error
	}
	return toCoreList(dataEvents), nil
}
