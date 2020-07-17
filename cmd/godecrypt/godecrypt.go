package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

func init(){
	flag.StringVar(&text, "t", "", "Text to decrypt")
	flag.StringVar(&passphrase, "p", "", "")
	flag.StringVar(&passphrase2, "p2", "", "")
	flag.Parse()
	if text == "" || passphrase == "" || passphrase2 == ""{
		fmt.Println("[GoDecrypt] Please provide passphrase 1, passphrase 2, and text to be decrypted by adding '-t \"<text>\"' and '-p <pass>' and '-p2 <pass>'")
		os.Exit(0)
	}
	fmt.Println("[GoEncrypt] Version 2.0")
}

var text,passphrase,passphrase2 string
var key []byte

func main(){
	key, _ := hex.DecodeString(passphrase)
	ciphertext, _ := hex.DecodeString(text)
	nonce, _ := hex.DecodeString(passphrase2)
	block, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(block)
	plaintext, _ := aesgcm.Open(nil, nonce, ciphertext, nil)
	fmt.Print("[GoDecrypt] Decrypted text: ")
	fmt.Printf("%s\n", plaintext)
}