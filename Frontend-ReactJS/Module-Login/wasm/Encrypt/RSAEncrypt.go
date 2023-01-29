package Encrypt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"sync"

	"golang.org/x/crypto/blake2b"
)

func Signature(PrivateKey string, Message []byte) (string, error) {
	privateKey := DecodeRSAPrivateKeyFromPemString(PrivateKey)
	random := rand.Reader
	hashed := sha256.Sum256(Message)
	signature, err := rsa.SignPKCS1v15(random, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(signature), nil
}

func EncryptRSA(privateKey string, publicKey string, secretMessage string) (string, error) {
	_publicKey := DecodeRSAPublicKeyFromPemString(publicKey)
	random := rand.Reader
	label := []byte("orders")
	hash, _ := blake2b.New(64, nil)
	ciphertext, err := rsa.EncryptOAEP(hash, random, _publicKey, []byte(secretMessage), label)
	if err != nil {
		return "", err
	}
	signature, err := Signature(privateKey, append([]byte(publicKey)[0:15], ciphertext[0:15]...))
	if err != nil {
		return "", err

	}
	return base64.URLEncoding.EncodeToString(ciphertext) + "." + signature, nil
}

func DecryptRSA(privateKey string, publicKey string, ciphertext string) (string, error) {
	_ciphertext := strings.Split(ciphertext, ".")
	_privateKey, _publicKey := DecodeRSAKeysFromPemString(privateKey, publicKey)
	var wg sync.WaitGroup
	temp_Ciphertext_ch := make(chan []byte)
	signature_ch := make(chan []byte)
	wg.Add(1)
	go func() {
		result, _ := base64.URLEncoding.DecodeString(_ciphertext[0])
		wg.Done()
		temp_Ciphertext_ch <- result
	}()
	wg.Add(1)
	go func() {
		result, _ := base64.URLEncoding.DecodeString(_ciphertext[1])
		wg.Done()
		signature_ch <- result
	}()
	wg.Wait()
	temp_Ciphertext, signature := <-temp_Ciphertext_ch, <-signature_ch

	hashed := sha256.Sum256(append([]byte(publicKey)[0:15], temp_Ciphertext[0:15]...))
	err := rsa.VerifyPKCS1v15(_publicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		return "", err
	}

	random := rand.Reader
	label := []byte("orders")
	hash, _ := blake2b.New(64, nil)
	plaintext, err := rsa.DecryptOAEP(hash, random, _privateKey, temp_Ciphertext, label)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(plaintext), nil
}
