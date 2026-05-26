package main

import (
	"fmt"
	"log"
	"time"

	"grpcdemo/cmd/server/examplepb"

	"github.com/getsentry/sentry-go"
	sentrygrpc "github.com/getsentry/sentry-go/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const grpcServerAddress = "localhost:50051"

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

	// Create a connection to the gRPC server with Sentry interceptors
	conn, err := grpc.NewClient(
		grpcServerAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()), // Use TLS in production
		grpc.WithUnaryInterceptor(sentrygrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(sentrygrpc.StreamClientInterceptor()),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %s", err)
	}
	defer conn.Close()

	// Create a client for the ExampleService
	client := examplepb.NewExampleServiceClient(conn)

	// Perform Unary call
	fmt.Println("Performing Unary Call:")
	unaryExample(client)

	// Perform Streaming call
	fmt.Println("\nPerforming Streaming Call:")
	streamExample(client)
}

func unaryExample(client examplepb.ExampleServiceClient) { _ = "STUB: not implemented"; return }

// Add metadata to the context

// Change to "error" to simulate an error

func streamExample(client examplepb.ExampleServiceClient) { _ = "STUB: not implemented"; return }

// Add metadata to the context

// Send multiple messages in the stream

// Close the stream for sending

// Receive responses from the server
