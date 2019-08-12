package main

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go-grpc-auth-demo/pb"
	"go-grpc-auth-demo/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gRPCServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(server.AuthFunc),
			server.AuthorizationUnaryServerInterceptor(),
		)),
	)

	pb.RegisterHelloServiceServer(gRPCServer, server.NewHelloServer())

	err = gRPCServer.Serve(listener)
	if err != nil {
		log.Fatalf("could not start gRPC server: %v", err)
	}
}
