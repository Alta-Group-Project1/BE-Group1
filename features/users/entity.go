package users

import "time"

type Core struct {
	ID          int
	FullName    string
	UserName    string
	Email       string
	Password    string
	PhoneNumber string
	Address     string
	ImageURL    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Business interface {
	InsertUser(data Core) (row int, err error)
	LoginUser(email string, password string) (userName string, token string, e error)
}

type Data interface {
	PostUser(data Core) (row int, err error)
	AuthUser(email string, password string) (userName string, token string, e error)
}
