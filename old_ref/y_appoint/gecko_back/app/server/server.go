// Package server is the core of the gRPC
// this will replace the router package
package server

import (
	"context"
	"gecko/auth"
	"gecko/crossLogging"
	"gecko/dbConn"
	"gecko/proto/pkg/authentication"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
)

type AuthServer struct {
	authentication.UnimplementedAuthenticationServer
}

func NewAuthServer() *AuthServer {
	return &AuthServer{}
}

func (a *AuthServer) Login(ctx context.Context, req *authentication.LoginRequest) (*authentication.LoginResponse, error) {

	token, err := auth.GetToken(ctx, req)
	if err != nil {
		crossLogging.Logger.Error().Err(err).Msg("failed to login")
		return nil, err
	}

	return token, nil

}

func Server() {
	err := godotenv.Load()
	if err != nil {
		crossLogging.Logger.Fatal().Stack().Err(err).Msg("failed to load env")
	}

	port := os.Getenv("GPORT")
	if port == "" {
		panic("PORT is not set. please set PORT")
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		crossLogging.Logger.Fatal().Err(err).Msg("failed to listen")
	}

	crossLogging.Logger.Info().Msg("lister started")
	server := grpc.NewServer()

	// register services here
	authentication.RegisterAuthenticationServer(server, NewAuthServer())

	reflection.Register(server)

	go func() {
		crossLogging.Logger.Info().Msg("grpc server is starting")

		if err := server.Serve(listener); err != nil {
			crossLogging.Logger.Fatal().Err(err).Msg("failed to serve grpc")
		}

		crossLogging.Logger.Info().Msg("grpc server is started")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	crossLogging.Logger.Info().Msg("grpc server is stopping")
	server.GracefulStop()

	crossLogging.Logger.Info().Msg("grpc server is stopped")
}

type AuthenticationServer interface {
	Login(ctx context.Context, req *authentication.LoginRequest) (*authentication.LoginResponse, error)
	mustEmbedUnimplementedGreetingServiceServer()
}
