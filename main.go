package main

import (
	"fmt"

	"service/environment"
	"service/server"
	"service/server/endpoints/index"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Sprintf("failed initialising zap logger: %s", err))
	}
	defer logger.Sync()

	// Environment variable defaulting and setting
	env, err := environment.Load()
	if err != nil {
		logger.Fatal("failed loading environment", zap.Error(err))
	}

	// Configure and setup the HTTP server with it's endpoints.
	config := &server.Config{
		Addr:   env.Address,
		Logger: logger,
	}

	s, err := server.New(
		config,
		&index.Endpoint{},
	)

	if err != nil {
		logger.Fatal("failed creating server", zap.Error(err))
	}

	logger.Info("starting server", zap.String("address", config.Addr))

	if err := s.Start(); err != nil {
		logger.Fatal("fatal server error", zap.Error(err))
	}
}
