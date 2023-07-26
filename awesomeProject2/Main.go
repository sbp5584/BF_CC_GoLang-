package main

import (
	"fmt"
)

var codebook = [4][2]int{{0b00, 0b01}, {0b01, 0b10}, {0b10, 0b11}, {0b11, 0b00}}
var message = [4]int{0b00, 0b01, 0b10, 0b11}
var cipherOFB = [4]int{}
var iv int = 0b10

func codebookLookup(XOR int) (lookupValue int) {
	var i, j int = 0, 0
	for i = 0; i < 4; i++ {
		if codebook[i][j] == XOR {
			j++
			lookupValue = codebook[i][j]
			break
		}
	}
	return lookupValue
}

func main() {
	var XOR int = 0
	var lookupValue int = 0
	lookupValue = codebookLookup(iv)

	for i := 0; i < 4; i++ {
		fmt.Printf("The plaintext value of is %02b\n", message[i])
	}

	//Ciphertext
	lookupValue = iv
	for i := 0; i < 4; i++ {
		XOR = message[i] ^ lookupValue
		lookupValue = codebookLookup(lookupValue)
		fmt.Printf("The OFB ciphered value of is %02b\n", XOR)
		cipherOFB[i] += XOR
	}

	//OFB Plaintext
	lookupValue = iv
	for i := 0; i < 4; i++ {
		XOR = cipherOFB[i] ^ lookupValue
		lookupValue = codebookLookup(lookupValue)
		fmt.Printf("The plaintext value of is %02b\n", XOR)
	}
}
