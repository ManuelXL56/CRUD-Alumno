package main

import (
	"fmt"
	"os"
	"runtime"

	"Module.com/Encription/Encrypt"
)

func init() {
	runtime.GOMAXPROCS(2)
}

func GenerateKeys_RSA_AES() string {
	var err error
	PrivateKey, PublicKey, err := Encrypt.GenerateKeysRSA(4096)
	if err != nil {
		panic(err)
	}
	Key, Nonce, err := Encrypt.GenerateKeyAES()
	if err != nil {
		panic(err)
	}

	rusult := "VITE_PrivateKey_RSA = " + PrivateKey + "\n" +
		"VITE_PublicKey_RSA = " + PublicKey + "\n" +
		"VITE_PASSWORD_AES_Key = " + Key + "\n" +
		"VITE_PASSWORD_AES_Key_Nonce = " + Nonce
	return rusult
}

func main() {
	// Create File
	file, err := os.Create("../.env.local")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	file.WriteString(GenerateKeys_RSA_AES())
	fmt.Println("Environment created successfully")
}
