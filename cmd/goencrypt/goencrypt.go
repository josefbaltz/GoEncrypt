package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

func init(){
	flag.StringVar(&text, "t", "", "Text to encrypt")
	flag.Parse()
	if text == "" {
		fmt.Println("[GoEncrypt] Please provide text to be encrypted by adding '-t <text>'")
		os.Exit(0)
	}
	fmt.Println("[GoEncrypt] Version 1.0")
}

var text string
var key []byte

func main(){
	bytes := make([]byte, 32)
	rand.Read(bytes)
	passphrase := hex.EncodeToString(bytes)
	fmt.Println("[GoEncrypt] Your Passphrase 1 is: " + passphrase)
	key, _ = hex.DecodeString(passphrase)
	plaintext := []byte(text)
	block, _ := aes.NewCipher(key)
	nonce := make([]byte, 12)
	rand.Read(nonce)
	noncehex := hex.EncodeToString(nonce)
	fmt.Println("[GoEncrypt] Your Passphrase 2 is: " + noncehex)
	aesgcm, _ := cipher.NewGCM(block)
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	fmt.Print("[GoEncrypt] Encrypted text: ")
	fmt.Printf("%x\n", ciphertext)
}