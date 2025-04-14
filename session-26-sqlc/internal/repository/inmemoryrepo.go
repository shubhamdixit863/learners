package repository

import (
	"context"
	"errors"
	"session-23-gin-jwt/internal/models"
)

type InMemory struct {
	users map[string]models.User
}

func (i *InMemory) UpdateUser(ctx context.Context, user models.User) error {
	//TODO implement me
	panic("implement me")
}

func (i *InMemory) DeleteUser(ctx context.Context, user models.User) error {
	//TODO implement me
	panic("implement me")
}

func (i *InMemory) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	var slc []*models.User
	for _, v := range i.users {
		slc = append(slc, &v)

	}

	return slc, nil

}

func (i *InMemory) GetUserByUserName(ctx context.Context, userName string) (*models.User, error) {
	// We will
	if v, ok := i.users[userName]; ok {

		return &v, nil
	}
	return nil, errors.New("user Not Found")
}

func (i *InMemory) CreateUser(ctx context.Context, user models.User) (string, error) {
	// We wil create the user here
	i.users[user.Username] = user

	return user.ID, nil
}

func NewInMemory() DbRepository {
	return &InMemory{
		users: make(map[string]models.User),
	}
}
