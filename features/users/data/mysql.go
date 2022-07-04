package data

import (
	_bcrypt "altaproject3/bcrypt"
	"altaproject3/features/users"
	"fmt"

	"altaproject3/middlewares"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(conn *gorm.DB) users.Data {
	return &mysqlUserRepository{
		db: conn,
	}
}

func (repo *mysqlUserRepository) PostUser(input users.Core) (row int, err error) {
	passHash, _ := _bcrypt.HashPassword(input.Password)
	user := User{
		FullName:    input.FullName,
		UserName:    input.UserName,
		Password:    passHash,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Address:     input.Address,
	}
	tx := repo.db.Create(&user)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to insert data")
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlUserRepository) AuthUser(email string, password string) (userName string, token string, e error) {
	userData := User{}
	repo.db.Where("email = ?", email).First(&userData)
	bool1 := _bcrypt.CheckPasswordHash(password, userData.Password)

	if !bool1 {
		return "", "", fmt.Errorf("error")
	}
	token, errToken := middlewares.CreateToken(int(userData.ID))
	if errToken != nil {
		return "", "", errToken
	}
	return token, userData.UserName, nil
}
