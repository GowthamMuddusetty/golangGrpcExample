package main

import (
	calculatorpb "chasistest/calculatortest"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := calculatorpb.NewCalculatorServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Add(ctx, &calculatorpb.AddRequest{A: 10, B: 100})
	if err != nil {
		log.Fatalf("Add: %v", err)
	}
	log.Printf("Greeting: %f", r.GetResult())

	subRes, err := c.Subtract(ctx, &calculatorpb.SubtractRequest{A: 10, B: 100})
	if err != nil {
		log.Fatalf("Substract: %v", err)
	}
	log.Printf("Greeting: %f", subRes.GetResult())

	mulResp, err := c.Multiply(ctx, &calculatorpb.MultiplyRequest{A: 10, B: 100})
	if err != nil {
		log.Fatalf("Multiply: %v", err)
	}
	log.Printf("Greeting: %f", mulResp.GetResult())

	divideRes, err := c.Divide(ctx, &calculatorpb.DivideRequest{A: 10, B: 100})
	if err != nil {
		log.Fatalf("Divide: %v", err)
	}
	log.Printf("Greeting: %f", divideRes.GetResult())
}
