package Encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// https://go.dev/src/crypto/cipher/example_test.go

func EncryptAES_GCM(key string, nonce string, text string) (string, error) {
	// Load your secret key from a safe place and reuse it across multiple
	// Seal/Open calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).
	SuperKey, _ := base64.URLEncoding.DecodeString(key)
	SuperNonce, _ := base64.URLEncoding.DecodeString(nonce)
	plaintext := []byte(text)

	block, err := aes.NewCipher(SuperKey)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := aesgcm.Seal(nil, SuperNonce, plaintext, nil)
	//fmt.Println("Len of ciphertext: ", len(ciphertext))
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func DecryptAES_GCM(key string, nonce string, ciphertext string) (string, error) {
	// Load your secret key from a safe place and reuse it across multiple
	// Seal/Open calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).
	SuperKey, _ := base64.URLEncoding.DecodeString(key)
	SuperNonce, _ := base64.URLEncoding.DecodeString(nonce)
	SuperCiphertext, _ := base64.URLEncoding.DecodeString(ciphertext)

	block, err := aes.NewCipher(SuperKey)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := aesgcm.Open(nil, SuperNonce, SuperCiphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), err
}
