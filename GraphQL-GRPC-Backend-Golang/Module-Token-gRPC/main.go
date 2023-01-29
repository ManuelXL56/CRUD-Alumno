package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"log"
	"net"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc"

	db "Module.com/JWT-Token/grpc-db-users"
	encrypt "Module.com/JWT-Token/grpc-encryption"
	token "Module.com/JWT-Token/grpc-token"
)

type server struct {
	token.UnimplementedRouteGuideServer
}

type MyCustomClaims struct {
	Matricule string
	Role      string
	jwt.RegisteredClaims
}

func (s *server) CreateToken(ctx context.Context, in *token.DataTokenRequest) (*token.Data, error) {
	//We get keys rsa and desecryption data login in format json
	KeysRSA, err := encrypt.GetKeysRSA(&encrypt.Empty{})
	if err != nil {
		return nil, err
	}
	DataLogin, err := encrypt.DecryptionRSA(&encrypt.DataRSA{PublicKey: in.GetPublicKey(), Data: in.GetData()})
	if err != nil {
		return nil, err
	}
	// Pass data login in format json to interface and execute and check login
	var JsonLogin *db.LoginRequest
	json.Unmarshal([]byte(DataLogin), &JsonLogin)
	// fmt.Println(DataLogin)
	GetLogin, err := db.Login(JsonLogin)
	if err != nil {
		return nil, err
	}
	if GetLogin.GetResult() != "true" {
		return nil, errors.New("valid token failed")
	}
	// We get private key rsa in their respective format
	PrivateKey, err := encrypt.DecodeRSAPrivateKeyFromPemString(&encrypt.KeysRSA{
		PrivateKey: KeysRSA.GetPrivateKey(),
	})
	if err != nil {
		return nil, err
	}
	ttl := (24 * time.Hour)
	// We get keys aes GCM and encrypted data user from token
	KeysAES, err := encrypt.GetKeysAES(&encrypt.Empty{})
	if err != nil {
		return nil, err
	}
	Matricule, err := encrypt.Encryption_AES(&encrypt.KeysAES_Data{Keys: &encrypt.KeysAES{Key: KeysAES.GetKey(), Nonce: KeysAES.GetNonce()}, Data: JsonLogin.GetMatricule()})
	if err != nil {
		return nil, err
	}
	Role, err := encrypt.Encryption_AES(&encrypt.KeysAES_Data{Keys: &encrypt.KeysAES{Key: KeysAES.GetKey(), Nonce: KeysAES.GetNonce()}, Data: GetLogin.GetRole()})
	if err != nil {
		return nil, err
	}
	// Create and signed token
	claims := MyCustomClaims{
		Matricule,
		Role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "Ritual Server",
			Subject:   "Ritual",
			Audience:  []string{"Ritual_else"},
		},
	}
	TokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	TokenString, err := TokenWithClaims.SignedString(PrivateKey)
	if err != nil {
		return nil, err
	}

	return &token.Data{
		Data: TokenString,
	}, nil
}

func (s *server) ValidateToken(ctx context.Context, in *token.DataTokenRequest) (*token.DataTokenReply, error) {
	// We generate keys rsa and we get the public key
	KeysRSA, err := encrypt.GetKeysRSA(&encrypt.Empty{})
	if err != nil {
		return nil, err
	}
	PublicKey, err := encrypt.DecodeRSAPublicKeyFromPemString(&encrypt.KeysRSA{
		PublicKey: KeysRSA.GetPublicKey(),
	})
	if err != nil {
		return nil, err
	}
	// We get the claims of the token with the public key rsa
	TokenWithClaims, err := jwt.ParseWithClaims(in.GetData(), &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return PublicKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := TokenWithClaims.Claims.(*MyCustomClaims); ok && TokenWithClaims.Valid {
		// We checking the validity token and get aes keys
		KeysAES, err := encrypt.GetKeysAES(&encrypt.Empty{})
		if err != nil {
			return nil, err
		}
		// We get keys aes GCM and encrypted data user from token
		Matricule, err := encrypt.Decryption_AES(&encrypt.KeysAES_Data{Keys: &encrypt.KeysAES{Key: KeysAES.GetKey(), Nonce: KeysAES.GetNonce()}, Data: claims.Matricule})
		if err != nil {
			return nil, err
		}
		Role, err := encrypt.Decryption_AES(&encrypt.KeysAES_Data{Keys: &encrypt.KeysAES{Key: KeysAES.GetKey(), Nonce: KeysAES.GetNonce()}, Data: claims.Role})
		if err != nil {
			return nil, err
		}
		//We search role from username and check if is your role
		RoleDB, err := db.SearchRole(Matricule)
		if RoleDB != Role || err != nil {
			return nil, errors.New("valid token failed")
		}
		// We encrypt the username that will send if the token and role are verified
		EscryptMatricule, err := encrypt.EncryptionRSA(&encrypt.DataRSA{PublicKey: in.GetPublicKey(), Data: Matricule})
		if err != nil {
			return nil, err
		}
		return &token.DataTokenReply{
			Matricule: EscryptMatricule,
			Result:    "valid token successfully",
		}, nil
	} else {
		return nil, errors.New("valid token failed")
	}
}

func main() {
	flag.Parse()
	port := ":50053"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	token.RegisterRouteGuideServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
