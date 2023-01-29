package Encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"sync"
)

func GenerateKeysRSA(sizeKey int) (Privatekey_pem string, PublicKey_pem string, err error) {
	random := rand.Reader
	PrivateKey, err := rsa.GenerateKey(random, sizeKey)
	if err != nil {
		return "", "", err
	}
	PublicKey := &PrivateKey.PublicKey

	var wg sync.WaitGroup
	PrivateKey_Ch := make(chan string)
	PublicKey_Ch := make(chan string)

	wg.Add(1)
	go GeneratePemPrivateKey(PrivateKey, PrivateKey_Ch, &wg)
	wg.Add(1)
	go GeneratePemPublicKey(PublicKey, PublicKey_Ch, &wg)

	wg.Wait()
	Privatekey_pem, PublicKey_pem = <-PrivateKey_Ch, <-PublicKey_Ch

	return
}
