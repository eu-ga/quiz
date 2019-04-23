package main

import (
	pb "github.com/eu-ga/quiz/proto"
	"github.com/eu-ga/quiz/server/service"
	"github.com/eu-ga/quiz/server/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":50051"
)

func main(){
	// Set-up gRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	//Create RPC service
	store := storage.DataStore{}
	// Register service with the gRPC server
	pb.RegisterQuizServer(s,&service.Service{store})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Println("Running on port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}