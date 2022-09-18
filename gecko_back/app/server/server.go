// Package server is the core of the gRPC
// this will replace the router package
package server

import (
	"database/sql"
	"gecko/crossLogging"
	"gecko/proto/pkg/authentication"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"net"
	"os"
	"os/signal"
)

type Repository interface{}

type db struct {
	db *bun.DB
}

func NewRepo() (Repository, error) {

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	open, err := sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+dbName)
	if err != nil {
		panic(err)
	}

	db := bun.NewDB(open, mysqldialect.New())

	return &db{db}, nil
}

type AuthServer struct {
	authentication.UnimplementedAuthenticationServer
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
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_opentracing.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(crossLogging.Logger),
			//grpc_auth.StreamServerInterceptor(),
			grpc_recovery.StreamServerInterceptor())),

		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(crossLogging.Logger),
			//grpc_auth.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
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
