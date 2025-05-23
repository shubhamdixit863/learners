package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-mysql-org/go-mysql/driver"
	"log"
	"time"
)

type User struct {
	FirstName  string
	SecondName string
}

func main() {
	// dsn format: "user:password@addr?dbname"
	dsn := "root:root@localhost:3306?users"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
		return
	}

	ctx, close := context.WithTimeout(context.Background(), 10*time.Second)
	defer close()
	conn, err := db.Conn(ctx)
	defer conn.Close()
	if err != nil {
		log.Println(err)
		return
	}

	//conn.ExecContext("")

	// conn
	queryContext, err := conn.QueryContext(ctx, "select FirstName,SecondName from Users")
	if err != nil {
		log.Println(err)
		return
	}
	var users []User

	i := 0
	for queryContext.Next() {
		user := User{}
		queryContext.Scan(&user.FirstName, &user.SecondName)
		i++
		users = append(users, user)
	}

	fmt.Println(users)

	// insert query
	//
	//result, err := conn.ExecContext(ctx, `INSERT  INTO Users(FirstName,SecondName,UserName,password) VALUES("shivam","gupta","some_username","some password")`)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	//fmt.Println(result.LastInsertId())
	//
	//fmt.Println(result.RowsAffected())

	//username := "some_username'); describe users;--"
	//query := fmt.Sprintf(`INSERT INTO Users (FirstName, SecondName, UserName, password) VALUES ("shivam", "gupta", "%s", "some password")`, username)
	//result, err := conn.ExecContext(ctx, query)
	//
	//if err != nil {
	//	log.Println(err)
	//}

	result, err := conn.ExecContext(
		ctx,
		`INSERT INTO Users (FirstName, SecondName, UserName, password) VALUES (?, ?, ?, ?)`,
		"shivam", "gupta", "some_username", "some password",
	)
	if err != nil {
		log.Println("Error inserting user:", err)
		return
	}

	log.Println(result)

	result, err = conn.ExecContext(
		ctx,
		`Update Users set FirstName=? where id=?`,
		"random", 1,
	)
	if err != nil {
		log.Println("Error inserting user:", err)
		return
	}

	result, err = conn.ExecContext(
		ctx,
		`Delete from Users  where id=?`,
		2,
	)
	if err != nil {
		log.Println("Error inserting user:", err)
		return
	}

	log.Println(result)

}
