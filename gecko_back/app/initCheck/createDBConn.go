package initCheck

import (
	"errors"
	"gecko/dbConn"
	"github.com/uptrace/bun"
)

func CreateDBConn() (*bun.DB, error) {
	db, err := dbConn.DBConnection()
	if err != nil {
		return nil, errors.New("failed to create db connection")
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.New("database does not correspond to this server")
	}

	return db, nil

}
