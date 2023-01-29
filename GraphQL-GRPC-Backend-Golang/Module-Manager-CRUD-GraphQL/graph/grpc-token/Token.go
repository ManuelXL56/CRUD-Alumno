package grpc_token

import (
	context "context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	encrypt "Module.com/GraphQL-Server/graph/grpc-encryption"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr_token *string = flag.String("addr_token", "svc-backend-grpc:50053", "the address to connect to")

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

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]

		KeysRSA, err := encrypt.GetKeysRSA(&encrypt.Empty{})
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		_, err = ValidateToken(&DataTokenRequest{PublicKey: KeysRSA.GetPublicKey(), Data: tokenString})
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.AbortWithStatus(http.StatusOK)
	}
}
