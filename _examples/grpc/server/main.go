package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/getsentry/sentry-go/_examples/grpc/server/examplepb"
	sentrygrpc "github.com/getsentry/sentry-go/grpc"
	"google.golang.org/grpc"
)

const grpcPort = ":50051"

// ExampleServiceServer is the server implementation for the ExampleService.
type ExampleServiceServer struct {
	examplepb.UnimplementedExampleServiceServer
}

// UnaryExample handles unary gRPC requests.
func (s *ExampleServiceServer) UnaryExample(ctx context.Context, req *examplepb.ExampleRequest) (*examplepb.ExampleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Simulate an error for demonstration

// StreamExample handles bidirectional streaming gRPC requests.
func (s *ExampleServiceServer) StreamExample(stream examplepb.ExampleService_StreamExampleServer) error {
	_ = "STUB: not implemented"
	return nil
}

func main() {
	// Initialize Sentry
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "",
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	// Create a new gRPC server with Sentry interceptors
	server := grpc.NewServer(
		grpc.UnaryInterceptor(sentrygrpc.UnaryServerInterceptor(sentrygrpc.ServerOptions{
			Repanic: true,
		})),
		grpc.StreamInterceptor(sentrygrpc.StreamServerInterceptor(sentrygrpc.ServerOptions{
			Repanic: true,
		})),
	)

	// Register the ExampleService
	examplepb.RegisterExampleServiceServer(server, &ExampleServiceServer{})

	// Start the server
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", grpcPort, err)
	}

	fmt.Printf("gRPC server is running on %s\n", grpcPort)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
