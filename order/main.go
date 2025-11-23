package main

import (
	"context"
	"log"
	"net"

	"github.com/sususaiasusubendera/common"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {
	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", grpcAddr, err)
	}
	defer l.Close()

	store := NewStore()
	service := NewService(store)
	NewGRPCHandler(grpcServer)

	service.CreateOrder(context.Background())

	log.Printf("Starting gRPC server at %s", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}
