package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"log"
	"session-23-gin-jwt/internal/models"
)

type MysqlRepo struct {
	conn *sqlx.DB
}

func (m MysqlRepo) UpdateUser(ctx context.Context, user models.User) error {
	//TODO implement me
	panic("implement me")
}

func (m MysqlRepo) DeleteUser(ctx context.Context, user models.User) error {
	//TODO implement me
	panic("implement me")
}

func (m MysqlRepo) CreateUser(ctx context.Context, user models.User) (interface{}, error) {

	result, err := m.conn.ExecContext(ctx, `INSERT INTO Users (FirstName, SecondName, UserName, password) VALUES (?, ?, ?, ?)`,
		user.FirstName, user.SeconName, user.Username, user.Password)

	if err != nil {
		log.Println(err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil

}

func (m MysqlRepo) GetUserByUserName(ctx context.Context, userName string) (*models.User, error) {
	var user models.User
	log.Println(userName)
	row := m.conn.QueryRowx("select Username,FirstName,SecondName from Users where userName=?", userName)
	err := row.Scan(&user.FirstName, &user.SeconName, &user.Username)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &user, nil
}

func (m MysqlRepo) GetAllUsers(ctx context.Context) ([]*models.User, error) {

	return nil, nil
}

func NewMysqlReqo(db *sqlx.DB) DbRepository {
	return &MysqlRepo{
		db,
	}

}
