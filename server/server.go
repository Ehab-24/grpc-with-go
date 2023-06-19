package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "suraj.com/rpc/inventory"
)

var (
	products []*pb.Product = make([]*pb.Product, 0)
	idCount  int32         = 1
)

func newID() int32 {
	defer func() { idCount++ }()
	return idCount
}

func findIndex(id *pb.ID) int {
	for i, p := range products {
		if p.Id == id.Val {
			return i
		}
	}
	return -1
}

func removeIndex(index int) {
	products = append(products[:index], products[index+1:]...)
}

type inventoryServer struct {
	pb.UnimplementedInventoryServer
}

func (s *inventoryServer) AddProduct(ctx context.Context, in *pb.Product) (*pb.AddProductResponse, error) {
	newId := newID()
	in.Id = newId

	products = append(products, in)

	return &pb.AddProductResponse{
		Acknowledged: true,
		InsertedId:   &pb.ID{Val: newId},
	}, nil

}

func (s *inventoryServer) GetProducts(ctx context.Context, in *pb.ProductsQuery) (*pb.ProductsResponse, error) {
	return &pb.ProductsResponse{
		Products: products,
	}, nil
}

func (s *inventoryServer) GetProduct(ctx context.Context, in *pb.ID) (*pb.Product, error) {
	if ind := findIndex(in); ind != -1 {
		return products[ind], nil
	}
	return nil, nil
}

func (s *inventoryServer) DeleteProduct(ctx context.Context, in *pb.ID) (*pb.DeleteProductResponse, error) {
	if ind := findIndex(in); ind != -1 {
		removeIndex(ind)
		return &pb.DeleteProductResponse{
			DeleteCount: 1,
		}, nil
	}
	return &pb.DeleteProductResponse{
		DeleteCount: 0,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterInventoryServer(grpcServer, &inventoryServer{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
