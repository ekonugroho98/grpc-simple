package grpc

import (
	"context"
	v1 "github.com/ekonugroho/grpc-simple/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

// RunServer runs gRPC service to publish Todo service
func RunServerGrpc(ctx context.Context, v1API v1.GreetServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// Register service
	server := grpc.NewServer()
	v1.RegisterGreetServiceServer(server, v1API)

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	go func() {
		for range c {
			log.Println("shutting down gRPC server...")
			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// Start gRPC server
	log.Println("starting gRPC server...")
	return server.Serve(listen)
}
