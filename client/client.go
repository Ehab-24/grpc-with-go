package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "suraj.com/rpc/inventory"
)

var (
	serverAddr = flag.String("server_addr", "localhost:50051", "The address of remote grpc server")
)

func main() {
	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewInventoryClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Add a new product
	newProduct := pb.Product{
		Name:        "Jelly Beans",
		Description: "lorem ipsum aut autem olis aetro",
		Quantity:    120,
		Price:       65,
	}
	client.AddProduct(ctx, &newProduct)

	// Get all products
	prods, err := client.GetProducts(ctx, &pb.ProductsQuery{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(prods)
}
