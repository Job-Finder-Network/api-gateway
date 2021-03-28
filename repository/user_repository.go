package repository

import (
	"context"
	"github.com/Job-Finder-Network/api-gateway/entity"
)

type UserRepository struct {
}

func (ur UserRepository) SelectUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	err := entity.DB.QueryRow("SELECT * FROM User where email=?", email).Scan(&user.ID, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (ur UserRepository) CreateUser(ctx context.Context, user entity.User) error {
	_, err := entity.DB.Query("insert into User(id,email,password,role) Values(?,?,?,?);", user.ID, user.Email, user.Password, user.Role)
	if err != nil {
		return err
	}
	return nil
}
