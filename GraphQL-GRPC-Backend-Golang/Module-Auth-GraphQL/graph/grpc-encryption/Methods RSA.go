package grpc_encryption

import (
	context "context"
	"flag"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr_encrypt_RSA *string = flag.String("addr_encrypt_RSA", "svc-backend-grpc:50051", "the address to connect to")

func GetKeysRSA(empty *Empty) (*KeysRSA, error) {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr_encrypt_RSA, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "did not connect: %s\n", err)
		return nil, err
	}
	defer conn.Close()
	client := NewRouteGuideClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := client.GetKeysRSA(ctx, empty)
	if err != nil {
		fmt.Fprintf(os.Stderr, "client.GenerateKeysRSA failed: %s\n", err)
		return nil, err
	}
	return result, nil
}
