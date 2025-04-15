package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// model which will be stored in Db
//type User struct {
//	Username  string `gorm:"column:UserName"`
//	Password  string `gorm:"column:password"`
//	FirstName string `gorm:"column:FirstName"`
//	SeconName string `gorm:"column:SecondName"`
//	ID        int    `gorm:"primaryKey"`
//}

//// TableName overrides the table name used by User to `profiles`
//func (User) TableName() string {
//	return "Users"
//}

type User struct {
	Username  string
	Password  string
	FirstName string
	SeconName string
	ID        primitive.ObjectID `bson:"_id,omitempty"`
}
