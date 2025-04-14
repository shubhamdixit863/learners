package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
	"session-23-gin-jwt/internal/models"
)

type MysqlOrm struct {
	db *gorm.DB
}

func (m MysqlOrm) CreateUser(ctx context.Context, user models.User) (int, error) {
	tx := m.db.Create(&user)
	if tx.Error != nil {
		log.Println("error", tx.Error)
		return 0, tx.Error
	}

	return user.ID, nil
}

func (m MysqlOrm) GetUserByUserName(ctx context.Context, userName string) (*models.User, error) {
	var user models.User
	log.Println("Incoming username", userName)
	tx := m.db.Where("UserName = ?", userName).Find(&user)
	if tx.Error != nil {
		log.Println("error", tx.Error)
		return nil, tx.Error
	}
	if user.Username == "" {
		return nil, errors.New("User not found")
	}
	return &user, nil
}

func (m MysqlOrm) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m MysqlOrm) UpdateUser(ctx context.Context, user models.User) error {
	//TODO implement me
	panic("implement me")
}

func (m MysqlOrm) DeleteUser(ctx context.Context, user models.User) error {
	//TODO implement me
	panic("implement me")
}

func NewMysqlOrm(db *gorm.DB) DbRepository {
	return &MysqlOrm{db: db}
}
