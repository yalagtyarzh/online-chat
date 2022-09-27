package database

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Conn struct {
	db *sql.DB
}

func New(user, pass, host, name string) (*Conn, error) {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?parseTime=true", user, pass, host, name)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}
	return &Conn{db: db}, nil
}

func (c *Conn) Close() error {
	return c.db.Close()
}
