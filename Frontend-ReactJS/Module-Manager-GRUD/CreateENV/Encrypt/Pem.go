package Encrypt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"sync"
)

func GeneratePemPrivateKey(PrivateKey *rsa.PrivateKey, ch chan<- string, wg *sync.WaitGroup) {
	_PrivateKey := x509.MarshalPKCS1PrivateKey(PrivateKey)
	wg.Done()
	ch <- base64.URLEncoding.EncodeToString(pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: _PrivateKey,
		},
	))
}

func GeneratePemPublicKey(PublicKey *rsa.PublicKey, ch chan<- string, wg *sync.WaitGroup) {
	_PublicKey := x509.MarshalPKCS1PublicKey(PublicKey)
	wg.Done()
	ch <- base64.URLEncoding.EncodeToString(pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: _PublicKey,
		},
	))
}
