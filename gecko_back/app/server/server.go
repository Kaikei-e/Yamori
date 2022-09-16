package server

import (
	"gecko/crossLogging"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
)

func Server() {
	port := os.Getenv("PORT")
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
