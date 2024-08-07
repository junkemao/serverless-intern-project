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
	pb.UnimplementedSpinServiceServer
}

func (s *server) Spin(ctx context.Context, req *pb.SpinRequest) (*pb.SpinResponse, error) {
	duration := time.Duration(req.DurationMs) * time.Millisecond

	start := time.Now()
	for {
		if time.Since(start) >= duration {
			break
		}
	}

	message := fmt.Sprintf("Busy spin loop completed after %v", duration)
	return &pb.SpinResponse{Message: message}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSpinServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
