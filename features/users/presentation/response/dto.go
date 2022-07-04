package response

import (
	"altaproject3/features/users"
	"time"
)

type User struct {
	ID          int       `json:"id"`
	FullName    string    `json:"full_name"`
	UserName    string    `json:"user_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	ImageURL    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
}

func FromCore(data users.Core) User {
	return User{
		ID:          data.ID,
		FullName:    data.FullName,
		UserName:    data.UserName,
		Email:       data.Email,
		PhoneNumber: data.PhoneNumber,
		Address:     data.Address,
		ImageURL:    data.ImageURL,
		CreatedAt:   data.CreatedAt,
	}
}

func FromCoreList(data []users.Core) []User {
	result := []User{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
