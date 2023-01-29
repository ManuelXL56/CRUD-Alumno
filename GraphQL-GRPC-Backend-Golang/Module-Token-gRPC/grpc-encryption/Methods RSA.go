package grpc_encryption

import (
	context "context"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
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

func EncryptionRSA(data *DataRSA) (string, error) {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr_encrypt_RSA, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "did not connect: %s\n", err)
		return "", err
	}
	defer conn.Close()
	client := NewRouteGuideClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := client.EncryptionRSA(ctx, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "client.EncryptionRSA failed: %s\n", err)
		return "", err
	}
	return result.GetData(), nil
}

func DecryptionRSA(data *DataRSA) (string, error) {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr_encrypt_RSA, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "did not connect: %s\n", err)
		return "", err
	}
	defer conn.Close()
	client := NewRouteGuideClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	Des_RSA, err := client.DecryptionRSA(ctx, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "client.DecryptionRSA failed: %s\n", err)
		return "", err
	}
	result, _ := base64.URLEncoding.DecodeString(Des_RSA.GetData())
	return string(result), nil
}

func DecodeRSAPrivateKeyFromPemString(KeysRSA *KeysRSA) (*rsa.PrivateKey, error) {
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

	result, err := client.DecodeRSAPrivateKeyFromPemString(ctx, KeysRSA)
	if err != nil {
		fmt.Fprintf(os.Stderr, "client.DecodeRSAPrivateKeyFromPemString failed: %s\n", err)
		return nil, err
	}

	GetPrivateKey, _ := base64.URLEncoding.DecodeString(result.GetPrivateKey())
	var PrivateKey *rsa.PrivateKey
	json.Unmarshal(GetPrivateKey, &PrivateKey)
	return PrivateKey, nil
}

func DecodeRSAPublicKeyFromPemString(KeysRSA *KeysRSA) (*rsa.PublicKey, error) {
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

	result, err := client.DecodeRSAPublicKeyFromPemString(ctx, KeysRSA)
	if err != nil {
		fmt.Fprintf(os.Stderr, "client.DecodeRSAPublicKeyFromPemString failed: %s\n", err)
		return nil, err
	}

	GetPublicKey, _ := base64.URLEncoding.DecodeString(result.GetPublicKey())
	var PublicKey *rsa.PublicKey
	json.Unmarshal(GetPublicKey, &PublicKey)
	return PublicKey, nil
}
