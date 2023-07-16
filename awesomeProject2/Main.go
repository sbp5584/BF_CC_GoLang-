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

	// ECB Mode
	fmt.Println("\nECB Mode:")
	ciphertextECB := make([]byte, len(plaintextBytes))
	for i := 0; i < len(plaintextBytes); i += aes.BlockSize {
		block.Encrypt(ciphertextECB[i:i+aes.BlockSize], plaintextBytes[i:i+aes.BlockSize])
	}
	fmt.Printf("Ciphertext: %x\n", ciphertextECB)

	// OFB Mode
	fmt.Println("\nOFB Mode:")
	ciphertextOFB := make([]byte, len(plaintextBytes))
	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(ciphertextOFB, plaintextBytes)
	fmt.Printf("Ciphertext: %x\n", ciphertextOFB)
}
