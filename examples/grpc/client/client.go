package main

import (
	"context"
	"go/examples/grpc/payment"
	"log"

	"google.golang.org/grpc"
)

const grpcServerAddress = "localhost:8000"

func main() {

	conn, err := grpc.Dial(grpcServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while making connection, %v", err)
	}
	defer conn.Close()

	c := payment.NewPaymentServiceClient(conn)

	tx := payment.TxRequest{
		From:   "0x123123123123",
		To:     "0x12313131",
		Amount: 50,
	}

	response, err := c.Transfer(context.Background(), &tx)
	if err != nil {
		log.Fatalf("Error when sending moeny %s", err)
	}

	log.Println(response)
}
