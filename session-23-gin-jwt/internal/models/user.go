package models

// model which will be stored in Db
type User struct {
	Username string
	Password string
	Name     string
	Email    string
	ID       string
}
