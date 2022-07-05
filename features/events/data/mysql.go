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

func (repo *mysqlEventRepository) SelectAllEvent(limit, offset int) ([]events.Core, error) {
	var dataEvents []Event
	result := repo.db.Preload("User").Limit(limit).Offset(offset).Find(&dataEvents)
	if result.Error != nil {
		return nil, result.Error
	}
	return toCoreList(dataEvents), nil
}

func (repo *mysqlEventRepository) InsertEvent(data events.Core) (int, error) {
	var dataEvent = fromCore(data)
	result := repo.db.Create(&dataEvent)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}
