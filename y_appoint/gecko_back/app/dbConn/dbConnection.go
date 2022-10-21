// Package dbConn is the connector to the database
// this package presents as a infrastructure layer of DDD
package dbConn

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"os"
)

type Conn interface {
	OpenConn() (*bun.DB, error)
	PassConn() (*bun.DB, error)
}

type DB struct {
	*bun.DB
}

func (db *DB) PassConn() (*bun.DB, error) {
	return db.DB, nil

}

func (db *DB) OpenConn() (*bun.DB, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	open, err := sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+dbName)
	if err != nil {
		panic(err)
	}

	db.DB = bun.NewDB(open, mysqldialect.New())

	return db.DB, nil
}
