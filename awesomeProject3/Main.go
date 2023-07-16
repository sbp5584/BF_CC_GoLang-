package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func main() {
	// Hardcoded key and IV
	key := []byte("0123456789ABCDEF")
	iv := []byte("ABCDEF1234567890")

	// Get plaintext from user
	var plaintext string
	fmt.Print("Enter plaintext: ")
	fmt.Scanln(&plaintext)

	// Convert plaintext to byte slice
	plaintextBytes := []byte(plaintext)

	// Pad plaintext to a multiple of the block size
	if len(plaintextBytes)%aes.BlockSize != 0 {
		padding := bytes.Repeat([]byte{0}, aes.BlockSize-len(plaintextBytes)%aes.BlockSize)
		plaintextBytes = append(plaintextBytes, padding...)
	}

	// Create new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// CBC Mode
	fmt.Println("\nCBC Mode:")
	ciphertextCBC := make([]byte, len(plaintextBytes))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertextCBC, plaintextBytes)
	fmt.Printf("Ciphertext: %x\n", ciphertextCBC)
}
