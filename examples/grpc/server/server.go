package main

import (
	"log"
	"net"

	"go/examples/grpc/payment"

	"google.golang.org/grpc"
)

func main() {
	port := ":8000"

	/// Specify the port we want to use to listen for client requests
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	/// Create an instance of the gRPC server
	gRPCServer := grpc.NewServer()

	/// Register the service with the newly created gRPC server
	payment.RegisterPaymentServiceServer(gRPCServer, &payment.PaymentServer{})

	/// Finally run server by listening to the port
	err = gRPCServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to start the gRPC server at port 8000 err %v", err)
	}

	log.Println("The server is running on  port 8000")
}
