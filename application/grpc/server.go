package grpc

import (
	"antrein/bc-dashboard/internal/pb"
	"antrein/bc-dashboard/model/config"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func StartServer(cfg *config.Config) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	helloService := &helloServer{}

	pb.RegisterGreeterServer(grpcServer, helloService)

	return grpcServer.Serve(lis)
}