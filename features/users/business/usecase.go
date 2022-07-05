package business

import (
	"altaproject3/features/users"

	"errors"
)

type userUsecase struct {
	userData users.Data
}

func NewUserBusiness(usrData users.Data) users.Business {
	return &userUsecase{
		userData: usrData,
	}
}

func (uc *userUsecase) InsertUser(input users.Core) (row int, err error) {
	if input.FullName == "" || input.UserName == "" || input.Password == "" || input.Email == "" || input.PhoneNumber == "" || input.Address == "" {
		return -1, errors.New("don't empty for all input")
	}
	row, err = uc.userData.PostUser(input)
	return row, err
}

func (uc *userUsecase) LoginUser(email string, password string) (userName string, token string, e error) {
	userName, token, e = uc.userData.AuthUser(email, password)
	return userName, token, e
}

func (uc *userUsecase) UpdateDataUser(id int, data users.Core) (row int, err error) {
	if data.FullName == "" || data.UserName == "" || data.Password == "" || data.Email == "" || data.PhoneNumber == "" || data.Address == "" || data.ImageURL == "" {
		return -1, errors.New("all input data must be filled")
	}
	row, err = uc.userData.PutDataUser(id, data)
	return row, err
}

func (uc *userUsecase) SelectUser(id int) (resp users.Core, err error) {
	resp, err = uc.userData.GetUser(id)
	return resp, err
}
