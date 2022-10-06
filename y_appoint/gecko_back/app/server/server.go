// Package server is the core of the gRPC
// this will replace the router package
package server

import (
	"context"
	"gecko/crossLogging"
	"gecko/dbConn"
	"gecko/proto/pkg/authentication"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpcopentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"net"
	"os"
	"os/signal"
)

type Repository interface{}

type AuthServer struct {
	authentication.UnimplementedAuthenticationServer
}

func (a *AuthServer) Login(ctx context.Context, req *authentication.LoginRequest) (*authentication.LoginResponse, error) {
	crossLogging.Logger.Info("login request", zap.String("username", req.Username))

	var conn dbConn.Conn
	db, err := conn.PassConn()
	if err != nil {
		crossLogging.Logger.Error("error while creating a new auth repo", zap.Error(err))
		return nil, err
	}

	user, err := dbConn.FetchUser(ctx, db, req.Username)
	if err != nil {
		crossLogging.Logger.Error("error while logging in", zap.Error(err))
		return nil, err
	}

	return &authentication.LoginResponse{
		Token: user.Token,
	}, nil

}

func NewAuthServer() *AuthServer {
	return &AuthServer{}
}

func Server() {
	err := godotenv.Load()
	if err != nil {
		crossLogging.Logger.Error("error while loading env", zap.Error(err))
	}

	port := os.Getenv("GPORT")
	if port == "" {
		panic("PORT is not set. please set PORT")
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		crossLogging.Logger.Error("error while listening", zap.Error(err))
	}

	crossLogging.Logger.Info("server is listening on port " + port)

	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpcctxtags.StreamServerInterceptor(),
			grpcopentracing.StreamServerInterceptor(),
			grpczap.StreamServerInterceptor(crossLogging.Logger),
			//grpc_auth.StreamServerInterceptor(),
			grpcrecovery.StreamServerInterceptor())),

		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpcctxtags.UnaryServerInterceptor(),
			grpcopentracing.UnaryServerInterceptor(),
			grpczap.UnaryServerInterceptor(crossLogging.Logger),
			//grpc_auth.UnaryServerInterceptor(),
			grpcrecovery.UnaryServerInterceptor(),
		)),
	)

	// register services here
	authentication.RegisterAuthenticationServer(server, NewAuthServer())

	reflection.Register(server)

	go func() {
		crossLogging.Logger.Info("grpc starting server")

		if err := server.Serve(listener); err != nil {
			crossLogging.Logger.Error("error while serving", zap.Error(err))
		}

		crossLogging.Logger.Info("grpc server was spawned")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	crossLogging.Logger.Info("grpc server shutting down")

	server.GracefulStop()

	crossLogging.Logger.Info("grpc server stopped")
}

type AuthenticationServer interface {
	Login(ctx context.Context, req *authentication.LoginRequest) (*authentication.LoginResponse, error)
	mustEmbedUnimplementedGreetingServiceServer()
}
