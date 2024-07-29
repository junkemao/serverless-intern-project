package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "demo/internal/pb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedExampleServiceServer
}

func (s *server) InvokeFunction(ctx context.Context, req *pb.FunctionRequest) (*pb.FunctionResponse, error) {
	defaultDuration := 5 * time.Second // Set the default duration for the busy spin loop

	start := time.Now()
	for {
		if time.Since(start) >= defaultDuration {
			break
		}
	}

	message := fmt.Sprintf("Busy spin loop completed after %v", defaultDuration)
	return &pb.FunctionResponse{Message: message}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterExampleServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
