package server

import (
	"context"
	"gecko/crossLogging"
	"gecko/proto/pkg/authentication"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

func NewAuthServer(db *bun.DB) *AuthServer {


	return &AuthServer{}
}

func (a *AuthServer) Login(context.Context, *authentication.LoginRequest) (*authentication.LoginResponse, error) {
	crossLogging.Logger.Info("login request", zap.String("username", req.Username))

	token, err :=
	if err != nil {
		crossLogging.Logger.Error("error while logging in", zap.Error(err))
		return nil, err
	}

	return token, nil

}
