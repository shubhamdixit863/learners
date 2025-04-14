package repository

import (
	"github.com/jmoiron/sqlx"
)

type MysqlDb struct {
	db *sqlx.Conn
}
