package main

import (
	"context"
	"log"
	"net"
    "time"

	"google.golang.org/grpc"
    "go.elastic.co/apm/module/apmgrpc"
    pb "github.com/charlottejhao/demo-apm/apm"
)

const (
    port = ":50051"
)

// server is used to implement apm.apmServer
type server struct{}

// Get implements apm.ApmServer
func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
    log.Printf("Received: simple request")
    return &pb.GetResponse{Reply: "Get simple response"}, nil
}

// GetDelay
func (s *server) GetDelay(ctx context.Context, in *pb.GetDelayRequest) (*pb.GetDelayResponse, error) {
    log.Printf("Received: delay request")

    time.Sleep(2 * time.Second)

     return &pb.GetDelayResponse{Reply: "Delay 2 sec response"}, nil
}

func main() {


	lis, err := net.Listen("tcp", port)

    if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(apmgrpc.NewUnaryServerInterceptor()))
	pb.RegisterApmServiceServer(s, &server{})

    if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
