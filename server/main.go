package main

import (
	"context"
	"log"
	"net"

	examplepb "chasistest/chasistest"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	examplepb.UnimplementedExampleServiceServer
}

func (s *Server) SayHello(ctx context.Context, req *examplepb.HelloRequest) (*examplepb.HelloResponse, error) {
	res := &examplepb.HelloResponse{
		Message: "Hello, " + req.GetName(),
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	examplepb.RegisterExampleServiceServer(s, &Server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
