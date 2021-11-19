package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

func main() {
	//crypto.Hash.String()
	b, _ := aes.NewCipher([]byte("Test1234Test1234"))
	data, err := ioutil.ReadFile("input.pdf")
	if err != nil {
		fmt.Println("Error :", err.Error())
	}
	fmt.Println("LEN:", len(data))
	gcm, err := cipher.NewGCM(b)
	if err != nil {
		log.Panic(err)
	}

	// Never use more than 2^32 random nonces with a given key
	// because of the risk of repeat.
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	// Save back to file
	err = ioutil.WriteFile("ciphertext.pdf", ciphertext, 0777)
	if err != nil {
		log.Panic(err)
	}

	//Decrypting

	reverseNonce := data[:gcm.NonceSize()]
	data = data[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, reverseNonce, data, nil)
	if err != nil {
		log.Panic(err)
	}

	err = ioutil.WriteFile("input.pdf", plaintext, 0777)
	if err != nil {
		log.Panic(err)
	}

}
