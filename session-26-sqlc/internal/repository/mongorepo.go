package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log"
	"session-23-gin-jwt/internal/models"
)

type MongoRepo struct {
	db *mongo.Client
}

func (m MongoRepo) CreateUser(ctx context.Context, user models.User) (interface{}, error) {
	coll := m.db.Database("users").Collection("user")
	log.Println("User here ----", user)
	result, err := coll.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	return result.InsertedID, nil
}

func (m MongoRepo) GetUserByUserName(ctx context.Context, userName string) (*models.User, error) {
	coll := m.db.Database("users").Collection("user")
	// Creates a query filter to match documents in which the "username" is

	filter := bson.D{{"username", userName}}
	// Retrieves the first matching document
	var result models.User
	err := coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (m MongoRepo) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m MongoRepo) UpdateUser(ctx context.Context, user models.User) error {
	//TODO implement me
	panic("implement me")
}

func (m MongoRepo) DeleteUser(ctx context.Context, user models.User) error {
	//TODO implement me
	panic("implement me")
}

func NewMongoRepo(db *mongo.Client) DbRepository {
	return &MongoRepo{db}

}
