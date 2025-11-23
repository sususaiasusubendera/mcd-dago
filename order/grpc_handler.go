package main

import (
	"context"
	"log"

	pb "github.com/sususaiasusubendera/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func NewGRPCHandler(grpcServer *grpc.Server) {
	handler := &grpcHandler{}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Printf("New order received! Order %v", req)
	o := &pb.Order{
		Id: "999",
	}
	return o, nil
}
