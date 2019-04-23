package main

import (
	pb "github.com/eu-ga/quiz/proto"
	"github.com/eu-ga/quiz/server/service"
	"github.com/eu-ga/quiz/server/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"math/rand"
	"net"
	"time"
)

var port = ""
var version = ""
var commit = ""
var buildTime = ""

func main() {
	log.Println("Version:", version)
	log.Println("Commit:", commit)
	log.Println("Build time:", buildTime)
	// Set-up gRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	//Create RPC service
	store := storage.DataStore{}
	// Register service with the gRPC server
	pb.RegisterQuizServer(s, &service.Service{store})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Println("Running on port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func init() {
	storage.Cache = storage.DataStore{
		Users:     make(map[int64]*storage.User),
		Questions: make(map[int64]*pb.Question),
	}
	rand.Seed(time.Now().UTC().UnixNano())
	storage.Cache.LoadQuestions()
}
