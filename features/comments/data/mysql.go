package data

import (
	"altaproject3/features/comments"
	"fmt"

	"gorm.io/gorm"
)

type mysqlCommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(conn *gorm.DB) comments.Data {
	return &mysqlCommentRepository{
		db: conn,
	}
}

func (repo *mysqlCommentRepository) InsertComment(input comments.Core) (int, error) {
	var dataComment = fromCore(input)
	comment := repo.db.Create(&dataComment)
	if comment.Error != nil {
		return 0, comment.Error
	}
	if comment.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to create comment")
	}
	return int(comment.RowsAffected), nil
}

func (repo *mysqlCommentRepository) DataGetAllComment(idEvent, limit, offset int) (data []comments.Core, err error) {
	var dataComment []Comment
	result := repo.db.Preload("User").Preload("Event").Where("event_id = ?", idEvent).Limit(limit).Offset(offset).Find(&dataComment)
	if result.Error != nil {
		return nil, result.Error
	}
	return toCoreList(dataComment), nil
}
