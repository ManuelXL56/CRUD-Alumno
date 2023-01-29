package grpc_token

import (
	context "context"
	"flag"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr_token *string = flag.String("addr_token", "svc-backend-grpc:50053", "the address to connect to")

func CreateToken(input *DataTokenRequest) (string, error) {
	count := 5
	GetLogin, err := BeforeCreateToken(input)
	for err != nil {
		if count == 0 {
			return "", err
		}
		GetLogin, err = BeforeCreateToken(input)
		count--
	}
	return GetLogin, nil
}

func BeforeCreateToken(input *DataTokenRequest) (string, error) {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr_token, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "did not connect: %s\n", err)
		return "", err
	}
	defer conn.Close()
	client := NewRouteGuideClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := client.CreateToken(ctx, input)
	if err != nil {
		//fmt.Fprintf(os.Stderr, "client.CreateToken failed: %s\n", err)
		return "", err
	}
	return result.GetData(), nil
}

func ValidateToken(input *DataTokenRequest) (*DataTokenReply, error) {
	count := 5
	GetLogin, err := BeforeValidateToken(input)
	for err != nil {
		if count == 0 {
			return nil, err
		}
		GetLogin, err = BeforeValidateToken(input)
		count--
	}
	return GetLogin, nil
}

func BeforeValidateToken(input *DataTokenRequest) (*DataTokenReply, error) {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr_token, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "did not connect: %s\n", err)
		return nil, err
	}
	defer conn.Close()
	client := NewRouteGuideClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := client.ValidateToken(ctx, input)
	if err != nil {
		//fmt.Fprintf(os.Stderr, "client.ValidateToken failed: %s\n", err)
		return nil, err
	}
	return result, nil
}
