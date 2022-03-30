package main

import (
	"fmt"
	"log"
	"net"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/avag-sargsyan/stricache/proto/stricache"
	"github.com/avag-sargsyan/stricache/internal/storage"
)

var (
	address string
)

func main() {
	opts := []grpc.ServerOption{
		grpc.MaxConcurrentStreams(200),
	}

	// create a gRPC server object
	grpcServer := grpc.NewServer(opts...)
	stricache.RegisterStricacheServiceServer(grpcServer, storage.NewCacheService())

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", "127.0.0.1:7999")
	if err != nil {
		log.Fatalf("Error in starting server %v", err)
	}
	fmt.Println("Started the server on:", address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("err in serving gRPC %v\n", err)
	}
}
