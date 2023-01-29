package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"runtime"
	"syscall/js"
	"wasm/Encrypt"
)

func init() {
	runtime.GOMAXPROCS(2)
}

func EncryptRSA() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		result, err := Encrypt.EncryptRSA(args[0].String(), args[1].String(), args[2].String())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error from: %s\n", err)
			return "Data Invalid"
		}
		return result
	})
}

func DecryptRSA() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		result, err := Encrypt.DecryptRSA(args[0].String(), args[1].String(), args[2].String())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error from: %s\n", err)
			return "Data Invalid"
		}
		Base64Decode, err := base64.URLEncoding.DecodeString(result)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error from: %s\n", err)
			return "Data Invalid"
		}
		return string(Base64Decode)
	})
}

func EncryptAES() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		result, err := Encrypt.EncryptAES_GCM(args[0].String(), args[1].String(), args[2].String())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error from: %s\n", err)
			return "Data Invalid"
		}
		return result
	})
}

func DecryptAES() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		result, err := Encrypt.DecryptAES_GCM(args[0].String(), args[1].String(), args[2].String())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error from: %s\n", err)
			return "Data Invalid"
		}
		return result
	})
}

func main() {
	ch := make(chan struct{})
	js.Global().Set("EncryptAES", EncryptAES())
	js.Global().Set("DecryptAES", DecryptAES())
	js.Global().Set("EncryptRSA", EncryptRSA())
	js.Global().Set("DecryptRSA", DecryptRSA())
	<-ch
}

/*func myGolangFunction() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return args[0].Int() + args[1].Int()
	})
}*/
