package dbConn

import (
	"gecko/crossLogging"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

type AuthRepo struct {
	db *bun.DB
}

func (a *AuthRepo) OpenConn() (*bun.DB, error) {

}

func NewAuthRepo() (DBConn, error) {
	if err != nil {
		crossLogging.Logger.Error("error while creating a new auth repo", zap.Error(err))
		return nil, err
	}

	return &AuthRepo{db: db}, nil
}

func (a *AuthRepo) Login(db DBConn, username, password string) (string, error) {
	db, err := NewAuthRepo()
	if err != nil {
		crossLogging.Logger.Error("error while creating a new auth repo", zap.Error(err))
		return "", err
	}

	return "", nil
}
