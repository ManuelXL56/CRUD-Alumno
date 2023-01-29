package Encrypt

import (
	"crypto/rand"
	"sync"
)

func GenerateRandomBytes_ch(n int, ch chan<- []byte, wg *sync.WaitGroup) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		panic(err)
	}
	wg.Done()
	ch <- b
}
