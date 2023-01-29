package grpc_db_users

import (
	context "context"
	"flag"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr_db *string = flag.String("addr_db", "svc-backend-grpc:50052", "the address to connect to")

func Login(DataLogin *LoginRequest) (*LoginReply, error) {
	count := 5
	result, err := BeforeLogin(DataLogin)
	for err != nil {
		if count == 0 {
			return nil, err
		}
		result, err = BeforeLogin(DataLogin)
		count--
	}
	return result, nil
}

func BeforeLogin(input *LoginRequest) (*LoginReply, error) {
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

	result, err := client.Login(ctx, input)
	if err != nil {
		//fmt.Fprintf(os.Stderr, "client.Login failed: %s\n", err)
		return nil, err
	}
	return result, nil
}

func SearchRole(Matricule string) (string, error) {
	count := 5
	result, err := BeforeSearchRole(&Data{Data: Matricule})
	for err != nil {
		if count == 0 {
			return "", err
		}
		result, err = BeforeSearchRole(&Data{Data: Matricule})
		count--
	}
	return result, nil
}

func BeforeSearchRole(input *Data) (string, error) {
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

	result, err := client.SearchUser(ctx, input)
	if err != nil {
		//fmt.Fprintf(os.Stderr, "client.SearchUser failed: %s\n", err)
		return "", err
	}
	return result.GetRole(), nil
}
