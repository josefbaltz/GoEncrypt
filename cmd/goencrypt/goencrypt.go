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

func main(){
	passbytes := make([]byte, 32) //Create array of 32 empty bytes
	rand.Read(passbytes) //Randomize data in previously created array
	passphrase := hex.EncodeToString(passbytes) //Encode the Bytes into a hexidecimal string
	fmt.Println("[GoEncrypt] Your Passphrase 1 is: " + passphrase) //Tell the user what that string is, known as 'Passphrase 1'
	plaintext := []byte(text) //Convert user provided plaintext into an array of bytes
	block, _ := aes.NewCipher(passbytes) //Create a new cipher block with 'passphrase 1'
	nonce := make([]byte, 12) //Create array of 12 empty bytes
	rand.Read(nonce) //Randomize data in previously created array
	noncehex := hex.EncodeToString(nonce) //Encode the nonce into a hexidecimal string
	fmt.Println("[GoEncrypt] Your Passphrase 2 is: " + noncehex) //Tell the user what that string is, known as 'Passphrase 2'
	aesgcm, _ := cipher.NewGCM(block) //Things I don't fully understand yet
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil) //Things I don't fully understand yet part 2 electric boogaloo
	fmt.Print("[GoEncrypt] Encrypted text: ") //Finally tell the user their encrypted text
	fmt.Printf("%x\n", ciphertext) //^
}