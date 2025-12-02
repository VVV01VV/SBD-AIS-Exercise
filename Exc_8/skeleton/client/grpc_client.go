package client

import (
	"context"
	"exc8/pb"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcClient struct {
	client pb.OrderServiceClient
}

func NewGrpcClient() (*GrpcClient, error) {
	conn, err := grpc.Dial(":4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewOrderServiceClient(conn)
	return &GrpcClient{client: client}, nil
}

func (c *GrpcClient) Run() error {
	ctx := context.Background()

	// 1. List drinks
	fmt.Println("Requesting drinks 🍹🍺☕")
	drinksResp, err := c.client.GetDrinks(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}

	fmt.Println("Available drinks:")
	for _, d := range drinksResp.Drinks {
		fmt.Printf("\t> id:%d  name:\"%s\"  price:%d  description:\"%s\"\n",
			d.Id, d.Name, d.Price, d.Description)
	}

	orderRound := func(qty int32, title string) error {
		fmt.Println(title)
		var items []*pb.OrderItem
		for _, d := range drinksResp.Drinks {
			fmt.Printf("\t> Ordering: %d x %s\n", qty, d.Name)
			items = append(items, &pb.OrderItem{
				DrinkId:  d.Id,
				Quantity: qty,
			})
		}
		_, err := c.client.OrderDrink(ctx, &pb.OrderDrinkRequest{Items: items})
		return err
	}
	// 2. Order a few drinks
	if err := orderRound(2, "Ordering drinks 👨‍🍳⏱️🍻🍻"); err != nil {
		return err
	}
	// 3. Order more drinks
	if err := orderRound(6, "Ordering another round of drinks 👨‍🍳⏱️🍻🍻"); err != nil {
		return err
	}
	// 4. Get order total
	fmt.Println("Getting the bill 💹💹💹")
	ordersResp, err := c.client.GetOrders(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}

	for _, t := range ordersResp.Totals {
		fmt.Printf("\t> Total: %d x %s\n", t.Quantity, t.Drink.Name)
	}

	return nil
}
