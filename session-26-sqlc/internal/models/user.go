package models

// model which will be stored in Db
type User struct {
	Username  string `gorm:"column:UserName"`
	Password  string `gorm:"column:password"`
	FirstName string `gorm:"column:FirstName"`
	SeconName string `gorm:"column:SecondName"`
	ID        int    `gorm:"primaryKey"`
}

// TableName overrides the table name used by User to `profiles`
func (User) TableName() string {
	return "Users"
}
