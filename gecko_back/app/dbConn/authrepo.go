package dbConn

import (
	"context"
	"gecko/crossLogging"
	"gecko/models/userrepo"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

func FetchUser(ctx context.Context, db *bun.DB, username string) (*userrepo.User, error) {
	crossLogging.Logger.Info("login request", zap.String("username", username))

	var user userrepo.User

	err := db.NewSelect().Model(&user).Where("username = ?", username).Scan(ctx)
	if err != nil {
		crossLogging.Logger.Error("error while creating a new auth repo", zap.Error(err))
		return nil, err
	}

	return &user, nil
}
