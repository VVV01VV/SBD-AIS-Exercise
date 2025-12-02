package server

import (
	"context"
	"exc8/pb"
	"net"

	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

type GRPCService struct {
	pb.UnimplementedOrderServiceServer

	drinks []*pb.Drink
	orders map[int32]int32 // drinkID -> total quantity
}

func StartGrpcServer() error {
	// Create a new gRPC server.
	srv := grpc.NewServer()
	// Create grpc service
	grpcService := &GRPCService{
		drinks: []*pb.Drink{
			{Id: 1, Name: "Spritzer", Price: 2, Description: "Wine with soda"},
			{Id: 2, Name: "Beer", Price: 3, Description: "Hagenberger Gold"},
			{Id: 3, Name: "Coffee", Price: 2, Description: "Mifare isn't that secure"},
		},
		orders: make(map[int32]int32),
	}
	// Register our service implementation with the gRPC server.
	pb.RegisterOrderServiceServer(srv, grpcService)
	// Serve gRPC server on port 4000.
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		return err
	}
	err = srv.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}

// todo implement functions
func (s *GRPCService) GetDrinks(ctx context.Context, _ *emptypb.Empty) (*pb.GetDrinksResponse, error) {
	return &pb.GetDrinksResponse{Drinks: s.drinks}, nil
}

func (s *GRPCService) OrderDrink(ctx context.Context, req *pb.OrderDrinkRequest) (*pb.OrderDrinkResponse, error) {
	for _, item := range req.Items {
		s.orders[item.DrinkId] += item.Quantity
	}
	return &pb.OrderDrinkResponse{
		Success: wrapperspb.Bool(true),
	}, nil
}

func (s *GRPCService) GetOrders(ctx context.Context, _ *emptypb.Empty) (*pb.GetOrdersResponse, error) {
	resp := &pb.GetOrdersResponse{}
	for id, qty := range s.orders {
		var drink *pb.Drink
		for _, d := range s.drinks {
			if d.Id == id {
				drink = d
				break
			}
		}
		if drink == nil {
			continue
		}
		resp.Totals = append(resp.Totals, &pb.OrderTotal{
			Drink:    drink,
			Quantity: qty,
		})
	}
	return resp, nil
}
