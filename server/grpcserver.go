package server

import (
	"fmt"
	"go-microservice/endpoint"
	"go-microservice/pb"
	"go-microservice/service"
	"net"
	"os"

	"github.com/go-kit/log"
	"google.golang.org/grpc"
)

func LaunchgRPCSever(port int) {
	var grpcSvc service.StringService

	grpcAddr := fmt.Sprintf(":%d", port)
	logger := log.NewLogfmtLogger(os.Stderr)

	grpcSvc = service.NewStringService()
	grpcSvc = service.NewLogging(logger, grpcSvc, "gRPC")
	endpoints := endpoint.MakeEndpoints(grpcSvc)

	var newGrpcServer pb.StringSvcServer
	newGrpcServer = endpoint.NewGrpcServer(endpoints)

	grpcListener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterStringSvcServer(baseServer, newGrpcServer)
		logger.Log("msg", "gRPC", "addr", grpcAddr)
		logger.Log("err", baseServer.Serve(grpcListener))
	}()
}
