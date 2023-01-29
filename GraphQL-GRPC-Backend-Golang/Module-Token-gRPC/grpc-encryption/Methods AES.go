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

var addr_encrypt_AES *string = flag.String("addr_encrypt_AES", "svc-backend-grpc:50051", "the address to connect to")

func GetKeysAES(empty *Empty) (*KeysAES, error) {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr_encrypt_AES, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "did not connect: %s\n", err)
		return nil, err
	}
	defer conn.Close()
	client := NewRouteGuideClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := client.GetKeysAES(ctx, empty)
	if err != nil {
		fmt.Fprintf(os.Stderr, "client.GenerateKeyAES failed: %s\n", err)
		return nil, err
	}
	return result, nil
}

func Encryption_AES(data *KeysAES_Data) (string, error) {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr_encrypt_AES, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "did not connect: %s\n", err)
		return "", err
	}
	defer conn.Close()
	client := NewRouteGuideClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := client.EncryptionAES(ctx, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "client.EncryptionAES failed: %s\n", err)
		return "", err
	}
	return result.GetData(), nil
}

func Decryption_AES(data *KeysAES_Data) (string, error) {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr_encrypt_AES, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "did not connect: %s\n", err)
		return "", err
	}
	defer conn.Close()
	client := NewRouteGuideClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := client.DecryptionAES(ctx, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "client.DecryptionAES failed: %s\n", err)
		return "", err
	}
	return result.GetData(), nil
}
