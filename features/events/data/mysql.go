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

func (repo *mysqlEventRepository) SelectEvent(idEvent int) (events.Core, error) {
	var dataEvent Event
	result := repo.db.Preload("User").Where("id = ?", idEvent).First(&dataEvent)
	if result.Error != nil {
		return events.Core{}, result.Error
	}
	return dataEvent.toCore(), nil
}

func (repo *mysqlEventRepository) InsertEvent(data events.Core) (int, error) {
	var dataEvent = fromCore(data)
	result := repo.db.Create(&dataEvent)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlEventRepository) DeleteEvent(idEvent int) (int, error) {
	result := repo.db.Where("id = ?", idEvent).Delete(&Event{})
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlEventRepository) UpdateEvent(idEvent int, data events.Core) (int, error) {
	var dataEvent = fromCore(data)
	result := repo.db.Model(&Event{}).Where("id = ?", idEvent).Updates(&dataEvent)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlEventRepository) SelectEventByUserId(idUser, limit, offset int) ([]events.Core, error) {
	var dataEvents []Event
	result := repo.db.Preload("User").Where("user_id = ? ", idUser).Limit(limit).Offset(offset).Find(&dataEvents)
	if result.Error != nil {
		return nil, result.Error
	}
	return toCoreList(dataEvents), nil
}

func (repo *mysqlEventRepository) CountEventData() (int, error) {
	var count int64
	result := repo.db.Model(&Event{}).Count(&count)
	if result.Error != nil {
		return -1, result.Error
	}
	return int(count), nil
}
