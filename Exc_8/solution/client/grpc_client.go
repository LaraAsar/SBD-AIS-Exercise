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
	conn, err := grpc.NewClient(":4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewOrderServiceClient(conn)
	return &GrpcClient{client: client}, nil
}

func (c *GrpcClient) Run() error {
	ctx := context.Background()

	// list drinks
	fmt.Println("Requesting drinks...")

	drinksResp, err := c.client.GetDrinks(ctx, &emptypb.Empty{})
	if err != nil {
		return fmt.Errorf("failed to get drinks: %v", err)
	}

	fmt.Println("Available drinks:")
	for _, d := range drinksResp.Drinks {
		fmt.Printf("\t> id:%d  name:\"%s\"  price:%d  description:\"%s\"\n",
			d.Id, d.Name, d.Price, d.Description)
	}

	// order drinks
	fmt.Println("Ordering drinks...")

	firstOrders := []pb.OrderItem{
		{DrinkId: 1, Quantity: 2},
		{DrinkId: 2, Quantity: 2},
		{DrinkId: 3, Quantity: 2},
	}

	for _, item := range firstOrders {
		fmt.Printf("\t> Ordering: %d x %s\n",
			item.Quantity, drinksResp.Drinks[item.DrinkId-1].Name)

		_, err := c.client.OrderDrink(ctx, &pb.OrderRequest{
			Item: &item,
		})
		if err != nil {
			return fmt.Errorf("failed to order drink: %v", err)
		}
	}

	// order more drinks
	fmt.Println("Ordering another round of drinks...")

	secondOrders := []pb.OrderItem{
		{DrinkId: 1, Quantity: 6},
		{DrinkId: 2, Quantity: 6},
		{DrinkId: 3, Quantity: 6},
	}

	for _, item := range secondOrders {
		fmt.Printf("\t> Ordering: %d x %s\n",
			item.Quantity, drinksResp.Drinks[item.DrinkId-1].Name)

		_, err := c.client.OrderDrink(ctx, &pb.OrderRequest{
			Item: &item,
		})
		if err != nil {
			return fmt.Errorf("failed to order drink: %v", err)
		}
	}

	// get order total
	fmt.Println("Getting the bill...")

	ordersResp, err := c.client.GetOrders(ctx, &emptypb.Empty{})
	if err != nil {
		return fmt.Errorf("failed to get orders: %v", err)
	}

	// aggregate totals per drink
	totals := map[int32]int32{}
	for _, order := range ordersResp.Orders {
		for _, item := range order.Items {
			totals[item.DrinkId] += item.Quantity
		}
	}

	// print totals like the exercise example
	for _, d := range drinksResp.Drinks {
		fmt.Printf("\t> Total: %d x %s\n", totals[d.Id], d.Name)
	}

	fmt.Println("Orders complete!")
	return nil
}
