package Encrypt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"log"
	"sync"
)

func DecodeRSAKeysFromPemString(PrivateKey string, PublicKey string) (*rsa.PrivateKey, *rsa.PublicKey) {
	var wg sync.WaitGroup
	PrivateKey_Ch := make(chan *rsa.PrivateKey)
	PublicKey_Ch := make(chan *rsa.PublicKey)

	wg.Add(1)
	go func() {
		result := DecodeRSAPrivateKeyFromPemString(PrivateKey)
		wg.Done()
		PrivateKey_Ch <- result
	}()
	wg.Add(1)
	go func() {
		result := DecodeRSAPublicKeyFromPemString(PublicKey)
		wg.Done()
		PublicKey_Ch <- result
	}()

	wg.Wait()
	_PrivateKey, _PublicKey := <-PrivateKey_Ch, <-PublicKey_Ch

	return _PrivateKey, _PublicKey
}

func DecodeRSAPrivateKeyFromPemString(PrivateKey string) *rsa.PrivateKey {
	GetPrivateKey, _ := base64.URLEncoding.DecodeString(PrivateKey)
	block, _ := pem.Decode(GetPrivateKey)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		log.Fatalf("failed to decode Private PEM block containing public key")
	}
	_PrivateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	return _PrivateKey
}

func DecodeRSAPublicKeyFromPemString(PublicKey string) *rsa.PublicKey {
	GetPublicKey, _ := base64.URLEncoding.DecodeString(PublicKey)
	block, _ := pem.Decode(GetPublicKey)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		log.Fatalf("failed to decode Public PEM block containing private key")
	}
	_PublicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	return _PublicKey
}
