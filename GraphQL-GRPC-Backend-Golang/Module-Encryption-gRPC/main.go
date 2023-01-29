package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"

	"google.golang.org/grpc"

	"Module.com/Encription/Encrypt"
	pb "Module.com/Encription/grpc-encryption"
)

type AES_Keys struct {
	Key   string
	Nonce string
}

type RSA_Keys struct {
	PrivateKey string
	PublicKey  string
}

var KeysRSA RSA_Keys
var KeysAES AES_Keys

func init() {
	runtime.GOMAXPROCS(2)

	var err error
	KeysRSA.PrivateKey, KeysRSA.PublicKey, err = Encrypt.GenerateKeysRSA(4096)
	if err != nil {
		panic(err)
	}
	KeysAES.Key, KeysAES.Nonce, err = Encrypt.GenerateKeyAES()
	if err != nil {
		panic(err)
	}
}

type server struct {
	pb.UnimplementedRouteGuideServer
}

func (s *server) GetKeysRSA(ctx context.Context, in *pb.Empty) (*pb.KeysRSA, error) {
	return &pb.KeysRSA{
		PrivateKey: KeysRSA.PrivateKey,
		PublicKey:  KeysRSA.PublicKey,
	}, nil
}

func (s *server) EncryptionRSA(ctx context.Context, in *pb.DataRSA) (*pb.Data, error) {
	result, err := Encrypt.EncryptRSA(KeysRSA.PrivateKey, in.GetPublicKey(), in.GetData())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from: %s\n", err)
		return nil, err
	}
	return &pb.Data{Data: result}, nil
}

func (s *server) DecryptionRSA(ctx context.Context, in *pb.DataRSA) (*pb.Data, error) {
	result, err := Encrypt.DecryptRSA(KeysRSA.PrivateKey, in.GetPublicKey(), in.GetData())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from: %s\n", err)
		return nil, err
	}
	return &pb.Data{Data: result}, nil
}

func (s *server) DecodeRSAKeysFromPemString(ctx context.Context, in *pb.KeysRSA) (*pb.KeysRSA, error) {
	PrivateKey, PublicKey := Encrypt.DecodeRSAKeysFromPemString(in.GetPrivateKey(), in.GetPublicKey())
	JsonPrivateKey, _ := json.Marshal(PrivateKey)
	JsonPublicKey, _ := json.Marshal(PublicKey)
	return &pb.KeysRSA{
		PrivateKey: base64.URLEncoding.EncodeToString(JsonPrivateKey),
		PublicKey:  base64.URLEncoding.EncodeToString(JsonPublicKey),
	}, nil
}

func (s *server) DecodeRSAPrivateKeyFromPemString(ctx context.Context, in *pb.KeysRSA) (*pb.KeysRSA, error) {
	PrivateKey := Encrypt.DecodeRSAPrivateKeyFromPemString(in.GetPrivateKey())
	JsonPrivateKey, _ := json.Marshal(PrivateKey)
	return &pb.KeysRSA{
		PrivateKey: base64.URLEncoding.EncodeToString(JsonPrivateKey),
	}, nil
}

func (s *server) DecodeRSAPublicKeyFromPemString(ctx context.Context, in *pb.KeysRSA) (*pb.KeysRSA, error) {
	PublicKey := Encrypt.DecodeRSAPublicKeyFromPemString(in.GetPublicKey())
	JsonPublicKey, _ := json.Marshal(PublicKey)
	return &pb.KeysRSA{
		PublicKey: base64.URLEncoding.EncodeToString(JsonPublicKey),
	}, nil
}

func (s *server) GetKeysAES(ctx context.Context, in *pb.Empty) (*pb.KeysAES, error) {
	return &pb.KeysAES{Key: KeysAES.Key, Nonce: KeysAES.Nonce}, nil
}

func (s *server) EncryptionAES(ctx context.Context, in *pb.KeysAES_Data) (*pb.Data, error) {
	result, err := Encrypt.EncryptAES_GCM(in.Keys.GetKey(), in.Keys.GetNonce(), in.GetData())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from: %s\n", err)
		return nil, err
	}
	return &pb.Data{Data: result}, nil
}

func (s *server) DecryptionAES(ctx context.Context, in *pb.KeysAES_Data) (*pb.Data, error) {
	result, err := Encrypt.DecryptAES_GCM(in.Keys.GetKey(), in.Keys.GetNonce(), in.GetData())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from: %s\n", err)
		return nil, err
	}
	return &pb.Data{Data: result}, nil
}

func main() {
	flag.Parse()
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRouteGuideServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
