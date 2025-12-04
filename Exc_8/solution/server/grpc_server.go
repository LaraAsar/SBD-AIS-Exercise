package server

import (
	"context"
	"exc8/pb"
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type GRPCService struct {
	pb.UnimplementedOrderServiceServer

	// In-memory storage for drinks and orders (required for assignment)
	drinks *pb.Drinks
	orders *pb.Orders
}

func StartGrpcServer() error {
	// Prepopulate drinks (as required by the assignment example)
	initialDrinks := &pb.Drinks{
		Drinks: []*pb.Drink{
			{Id: 1, Name: "Spritzer", Price: 2, Description: "Wine with soda"},
			{Id: 2, Name: "Beer", Price: 3, Description: "Hagenberger Gold"},
			{Id: 3, Name: "Coffee", Price: 0, Description: "Mifare isn't that secure"},
		},
	}

	// Prepare empty orders list
	initialOrders := &pb.Orders{
		Orders: []*pb.Order{},
	}

	// Create grpc service
	grpcService := &GRPCService{
		drinks: initialDrinks,
		orders: initialOrders,
	}

	// Create a new gRPC server.
	srv := grpc.NewServer()

	// Register our service implementation with the gRPC server.
	pb.RegisterOrderServiceServer(srv, grpcService)

	// Serve gRPC server on port 4000.
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		return err
	}

	slog.Info("gRPC server running on :4000")

	err = srv.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}

// todo implement functions

// GetDrinks returns the available drink list
func (s *GRPCService) GetDrinks(ctx context.Context, _ *emptypb.Empty) (*pb.Drinks, error) {
	return s.drinks, nil
}

// OrderDrink stores each incoming order as a new order entry
func (s *GRPCService) OrderDrink(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
	if req == nil || req.Item == nil {
		return &pb.OrderResponse{Success: false}, fmt.Errorf("invalid request")
	}

	// Create a new order with incremental ID
	newOrder := &pb.Order{
		Id:    int32(len(s.orders.Orders) + 1),
		Items: []*pb.OrderItem{req.Item},
	}

	// Save to in-memory list
	s.orders.Orders = append(s.orders.Orders, newOrder)

	return &pb.OrderResponse{Success: true}, nil
}

// GetOrders returns all previously stored orders
func (s *GRPCService) GetOrders(ctx context.Context, _ *emptypb.Empty) (*pb.Orders, error) {
	return s.orders, nil
}
