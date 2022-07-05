package business

import (
	"altaproject3/features/comments"
	"errors"
)

type commentUsecase struct {
	commentData comments.Data
}

func NewCommentBusiness(cmntData comments.Data) comments.Business {
	return &commentUsecase{
		commentData: cmntData,
	}
}

func (cm *commentUsecase) InsertComment(data comments.Core) (row int, err error) {
	if data.Content == "" {
		return -1, errors.New("comment cannot be empty")
	}
	row, err = cm.commentData.InsertComment(data)
	return row, err
}
