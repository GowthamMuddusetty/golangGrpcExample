package main

import (
	"context"
	"log"
	"net"

	calculatorpb "chasistest/calculatortest"

	"google.golang.org/grpc"
)

type CalServer struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (c *CalServer) Add(ctx context.Context, req *calculatorpb.AddRequest) (*calculatorpb.AddResponse, error) {
	val := &calculatorpb.AddResponse{
		Result: req.A + req.B,
	}
	return val, nil
}

func (c *CalServer) Subtract(ctx context.Context, req *calculatorpb.SubtractRequest) (*calculatorpb.SubtractResponse, error) {
	val := &calculatorpb.SubtractResponse{
		Result: req.A - req.B,
	}
	return val, nil
}

func (c *CalServer) Multiply(ctx context.Context, req *calculatorpb.MultiplyRequest) (*calculatorpb.MultiplyResponse, error) {
	val := &calculatorpb.MultiplyResponse{
		Result: req.A * req.B,
	}
	return val, nil
}

func (c *CalServer) Divide(ctx context.Context, req *calculatorpb.DivideRequest) (*calculatorpb.DivideResponse, error) {
	val := &calculatorpb.DivideResponse{
		Result: req.A / req.B,
	}
	return val, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &CalServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
