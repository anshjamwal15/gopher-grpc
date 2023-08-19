package main

import (
	"log"
	"net"

	"github.com/aimbot1526/test-go/db"
	pb "github.com/aimbot1526/test-go/generated"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	log.Printf("Server Running on port 8080")
	if err != nil {
		panic(err)
	}
	db.InitMongo()
	s := grpc.NewServer()

	inject := NewInject()

	pb.RegisterShopServiceServer(s, inject.ShopService)
	pb.RegisterProductServiceServer(s, inject.ProductService)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
