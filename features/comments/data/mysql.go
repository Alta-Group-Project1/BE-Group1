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
