package repository

import (
	"appoint/initialize"
	"database/sql"
	"fmt"
	mysql "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"os"
	"time"
)

type (
	DBConn interface {
		Connect() *bun.DB
		Ping() error
	}

	DBConnImpl struct {
	}
)

func (d *DBConnImpl) Connect() *bun.DB {
	envs := initialize.LoadEnv()

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		// エラーハンドリング
	}
	c := mysql.Config{
		DBName:               envs["dbName"],
		User:                 envs["dbUser"],
		Passwd:               envs["dbPass"],
		Addr:                 envs["dbIP"] + ":" + envs["dbPort"],
		Net:                  "tcp",
		ParseTime:            true,
		Collation:            "utf8mb4_unicode_ci",
		Loc:                  jst,
		AllowNativePasswords: true,
	}
	open, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error!! failed to initialize db : %+v", err)
	}

	db := bun.NewDB(open, mysqldialect.New())

	return db

}

func (d *DBConnImpl) Ping() error {
	db := d.Connect()
	raw := db.DB

	// TODO: error handling. shift to logger
	err := raw.Ping()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error!! failed to ping db : %+v", err)
		panic(err)
	}

	return nil
}

func NewDBConn() DBConn {
	return &DBConnImpl{}
}
