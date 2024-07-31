package main

import (
	"context"
	"log"
	"time"

	pb "demo/internal/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("grpc-server.default.192.168.1.240.sslip.io:80", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewExampleServiceClient(conn)

	rate := 1
	interval := time.Second / time.Duration(rate)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			log.Println("Client context done")
			return
		case <-ticker.C:
			req := &pb.FunctionRequest{} // Empty request
			res, err := client.InvokeFunction(ctx, req)
			if err != nil {
				log.Fatalf("could not invoke: %v", err)
			}
			log.Printf("Response: %s", res.Message)
		}
	}
}
