package main

import (
	"authorization/handler"
	pb "authorization/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("authorization"),
	)

	// Register handler
	if err := pb.RegisterAuthorizationHandler(srv.Server(), handler.New()); err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
