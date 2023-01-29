package grpc_db_users

import (
	context "context"
	"flag"
	"fmt"
	"os"
	"time"

	//pb "Module.com/GraphQL-Server/graph/grpc-database-users"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr_db *string = flag.String("addr_db", "svc-backend-grpc:50052", "the address to connect to")

func SearchUser(input *Data) (*DataUser, error) {
	count := 5
	result, err := BeforeSearchUser(input)
	for err != nil {
		if count == 0 {
			return nil, err
		}
		result, err = BeforeSearchUser(input)
		count--
	}
	return result, nil
}

func BeforeSearchUser(input *Data) (*DataUser, error) {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr_db, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "did not connect: %s\n", err)
		return nil, err
	}
	defer conn.Close()
	client := NewRouteGuideClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := client.SearchUser(ctx, input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func RegisterUser(input *DataUser) (string, error) {
	count := 5
	result, err := BeforeRegisterUser(input)
	for err != nil {
		if count == 0 {
			return "", err
		}
		result, err = BeforeRegisterUser(input)
		count--
	}
	return result, nil
}

func BeforeRegisterUser(input *DataUser) (string, error) {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr_db, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "did not connect: %s\n", err)
		return "", err
	}
	defer conn.Close()
	client := NewRouteGuideClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := client.RegisterUser(ctx, input)
	if err != nil {
		return "", err
	}
	return result.GetData(), nil
}

func UpdateUser(input *DataUser_Update) (string, error) {
	count := 5
	result, err := BeforeUpdateUser(input)
	for err != nil {
		if count == 0 {
			return "", err
		}
		result, err = BeforeUpdateUser(input)
		count--
	}
	return result, nil
}

func BeforeUpdateUser(input *DataUser_Update) (string, error) {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr_db, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "did not connect: %s\n", err)
		return "", err
	}
	defer conn.Close()
	client := NewRouteGuideClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := client.UpdateUser(ctx, input)
	if err != nil {
		return "", err
	}
	return result.GetData(), nil
}

func DeleteUser(input *Data) (string, error) {
	count := 5
	result, err := BeforeDeleteUser(input)
	for err != nil {
		if count == 0 {
			return "", err
		}
		result, err = BeforeDeleteUser(input)
		count--
	}
	return result, nil
}

func BeforeDeleteUser(input *Data) (string, error) {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr_db, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "did not connect: %s\n", err)
		return "", err
	}
	defer conn.Close()
	client := NewRouteGuideClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := client.DeleteUser(ctx, input)
	if err != nil {
		return "", err
	}
	return result.GetData(), nil
}
