package Encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"sync"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/scrypt"
)

// https://go.dev/src/crypto/cipher/example_test.go

// AES GCM

func GenerateKeyAES() (KeyAES string, nonce string, err error) {
	var wg sync.WaitGroup
	KEY_128Bits_Ch := make(chan []byte)
	KEY_512Bits_Ch := make(chan []byte)

	wg.Add(1)
	go GenerateRandomBytes_ch(32, KEY_128Bits_Ch, &wg)
	wg.Add(1)
	go GenerateRandomBytes_ch(64, KEY_512Bits_Ch, &wg)

	wg.Wait()
	KEY_128Bits, KEY_512Bits := <-KEY_128Bits_Ch, <-KEY_512Bits_Ch

	var salt []byte
	_blake2b, _ := blake2b.New512(KEY_512Bits)
	salt = _blake2b.Sum(KEY_512Bits)

	KeySuper, err := scrypt.Key(KEY_128Bits, salt, 1<<15, 8, 1, 32)
	if err != nil {
		return "", "", err
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	_nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, _nonce); err != nil {
		return "", "", err
	}

	KeyAES = base64.URLEncoding.EncodeToString(KeySuper)
	nonce = base64.URLEncoding.EncodeToString(_nonce)

	return KeyAES, nonce, nil
}

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
