package authrepo

import (
	"gecko/crossLogging"
	"gecko/dbConn"
	"go.uber.org/zap"
)

type AuthRepo struct {
	db dbConn.DBConn
}

func NewAuthRepo(db dbConn.DBConn) (*AuthRepo, error) {
	return &AuthRepo{db: db}, nil
}

func (a *AuthRepo) Login(db dbConn.DBConn, username, password string) (string, error) {
	db, err := NewAuthRepo(a.db)
	if err != nil {
		crossLogging.Logger.Error("error while creating a new auth repo", zap.Error(err))
		return "", err
	}

	return "", nil
}
