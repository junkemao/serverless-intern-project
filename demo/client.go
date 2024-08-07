package main

import (
	"bufio"
	"os"
	"context"
	"log"
	"time"

	pb "demo/internal/pb"

	"google.golang.org/grpc"
)

func main() {
	rate := 35
	interval := time.Second / time.Duration(rate)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()


	timeout := time.After(300 * time.Second)
	
	for {
		select {
		case <-timeout:
			log.Println("Finished")
			return
		case <-ticker.C:
			start := time.Now()
			conn, err := grpc.Dial("grpc-server-mod.default.192.168.1.240.sslip.io:80", grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()
			client := pb.NewSpinServiceClient(conn)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
			defer cancel()

			durationMs := int64(400)
			req := &pb.SpinRequest{DurationMs: durationMs}
			res, rerr := client.Spin(ctx, req)
			if rerr != nil {
				log.Fatalf("could not invoke: %v", err)
			}
			log.Printf("Response: %s", res.Message)
			end := time.Now()
			invoke_delay := end.Sub(start) - time.Millisecond * 400
			
			fo, _ := os.OpenFile("output.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

			w := bufio.NewWriter(fo)
			w.WriteString(invoke_delay.String() +"\n")

			w.Flush()
			fo.Close()

		}
	}

}
