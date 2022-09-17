// Package server is the core of the gRPC
// this will replace the router package
package server

import (
	"context"
	"gecko/crossLogging"
	"gecko/proto/pkg/authentication"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/grpc-ecosystem/go-grpc-middleware"

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

	token, err := a.Login(ctx, req)
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
	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_opentracing.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(zapLogger),
			grpc_auth.StreamServerInterceptor(myAuthFunction),
			grpc_recovery.StreamServerInterceptor()),

			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_ctxtags.UnaryServerInterceptor(),
				grpc_opentracing.UnaryServerInterceptor(),
				grpc_zap.UnaryServerInterceptor(zapLogger),
				grpc_auth.UnaryServerInterceptor(myAuthFunction),
				grpc_recovery.UnaryServerInterceptor(),
			)),
		)

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
