package repository

import (
	"errors"
	"session-23-gin-jwt/internal/models"
)

type InMemory struct {
	users map[string]models.User
}

func (i *InMemory) GetAllUsers() ([]*models.User, error) {
	var slc []*models.User
	for _, v := range i.users {
		slc = append(slc, &v)

	}

	return slc, nil

}

func (i *InMemory) GetUserByUserName(userName string) (*models.User, error) {
	// We will
	if v, ok := i.users[userName]; ok {

		return &v, nil
	}
	return nil, errors.New("user Not Found")
}

func (i *InMemory) CreateUser(user models.User) string {
	// We wil create the user here
	i.users[user.Username] = user

	return user.ID
}

func NewInMemory() DbRepository {
	return &InMemory{
		users: make(map[string]models.User),
	}
}
