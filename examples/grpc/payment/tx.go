package payment

import (
	context "context"
	"log"
)

type PaymentServer struct{}

// Handler
//ctx is used by the goroutines to interact with GRPC
func (ps *PaymentServer) Transfer(ctx context.Context, req *TxRequest) (*TxResponse, error) {
	log.Println("Sent $", req.Amount, "from ", req.From, "to ", req.To)
	return &TxResponse{Success: true}, nil
}
