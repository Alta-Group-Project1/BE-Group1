package business

import (
	"altaproject3/features/comments"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockCommentDataFailed struct{}

func (mock mockCommentDataFailed) InsertComment(data comments.Core) (row int, err error) {
	return 0, fmt.Errorf("failed to insert comment")
}

func (mock mockCommentDataFailed) DataGetAllComment(idEvent, limit, offset int) (data []comments.Core, err error) {
	return nil, fmt.Errorf("failed to get comment")
}

type mockCommentData struct{}

func (mock mockCommentData) InsertComment(data comments.Core) (row int, err error) {
	return 1, nil
}
func (mock mockCommentData) DataGetAllComment(idEvent, limit, offset int) (data []comments.Core, err error) {
	return []comments.Core{
		{ID: 1, Content: "alta",
			User: comments.User{
				ID: 1, UserName: "alta1",
				FullName: "alta-alterra",
				ImageURL: "axaswaw"},
			Event: comments.Event{
				ID: 1, EventName: "alta2",
				Price:      10000,
				Address:    "andalas",
				DateStart:  "1",
				DateFinish: "2"},
		},
	}, nil
}

func TestInsertComment(t *testing.T) {
	t.Run("check input comment succes", func(t *testing.T) {
		commentBusiness := NewCommentBusiness(mockCommentData{})
		newComment := comments.Core{
			Content: "alta",
			Event: comments.Event{
				ID: 1,
			},
		}
		result, err := commentBusiness.InsertComment(newComment)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})

	t.Run("check input comment failed", func(t *testing.T) {
		commentBusiness := NewCommentBusiness(mockCommentDataFailed{})
		newComment := comments.Core{
			Content: "alta",
			Event: comments.Event{
				ID: 1,
			},
		}
		result, err := commentBusiness.InsertComment(newComment)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})

	t.Run("check input comment failed when content empty", func(t *testing.T) {
		commentBusiness := NewCommentBusiness(mockCommentDataFailed{})
		newComment := comments.Core{
			Content: "",
			Event: comments.Event{
				ID: 1,
			},
		}
		result, err := commentBusiness.InsertComment(newComment)
		assert.NotNil(t, err)
		assert.Equal(t, -1, result)
	})
}

func TestGetAllComment(t *testing.T) {
	t.Run("Test Get All Comment Success", func(t *testing.T) {
		commentBusiness := NewCommentBusiness(mockCommentData{})
		result, err := commentBusiness.GetAllComment(0, 0, 0)
		assert.Nil(t, err)
		assert.Equal(t, "alta", result[0].Content)
	})

	t.Run("Test Get All Comment Failed", func(t *testing.T) {
		commentBusiness := NewCommentBusiness(mockCommentDataFailed{})
		result, err := commentBusiness.GetAllComment(0, 0, 0)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}
