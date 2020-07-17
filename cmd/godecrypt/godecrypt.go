package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func init(){
	flag.StringVar(&text, "t", "", "Text to decrypt")
	flag.StringVar(&input, "i", "", "File to decrypt")
	flag.StringVar(&passphrase, "p", "", "")
	flag.StringVar(&passphrase2, "p2", "", "")
	flag.Parse()
	if (text == "" && input == "") || passphrase == "" || passphrase2 == "" {
		fmt.Println("[GoDecrypt] Please provide passphrase 1, passphrase 2, and either text or file to be decrypted by adding '-t \"<text>\"' or -i <file location> and '-p <pass>' and '-p2 <pass>'")
		os.Exit(0)
	}
	fmt.Println("[GoDecrypt] Version 1.1")
}

var text,input,passphrase,passphrase2 string
var key,ciphertext []byte
var err error

func main(){
	if input == "" {
		ciphertext, _ = hex.DecodeString(text) //Decode hexadecimal ciphertext into bytes
	} else {
		ciphertext, err = ioutil.ReadFile(input)
		if err != nil {
			fmt.Println("[GoDecrypt] Error!\n"+err.Error())
		}
	}
	key, _ := hex.DecodeString(passphrase) //Decode hexadecimal passphrase 1 into bytes
	nonce, _ := hex.DecodeString(passphrase2) //Decode hexadecimal passphrase 2 into bytes
	block, _ := aes.NewCipher(key) //Create a new cipher block with decoded passphrase 1
	aesgcm, _ := cipher.NewGCM(block) //Things I don't fully understand yet
	plaintext, _ := aesgcm.Open(nil, nonce, ciphertext, nil) //Things I don't fully understand yet part 2 electric boogaloo
	if input == "" {
		fmt.Print("[GoDecrypt] Decrypted text: ") //Finally tell the user their decrypted text
		fmt.Printf("%s\n", plaintext)             //^
	} else {
		ioutil.WriteFile("./output.dec", plaintext, 0644)
	}
}