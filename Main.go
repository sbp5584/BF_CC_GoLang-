package main

import (
	"fmt"
)

func bruteForceCaesarCipher(s string) {
	for shift := 0; shift < 26; shift++ {
		decoded := []rune(s)

		for i, char := range decoded {
			if 'a' <= char && char <= 'z' {
				decoded[i] = rune((int(char-'a')+shift)%26 + 'a')
			} else if 'A' <= char && char <= 'Z' {
				decoded[i] = rune((int(char-'A')+shift)%26 + 'A')
			} else if char == ' ' {
				decoded[i] = ' '
			}
		}

		fmt.Printf("Shift %d: %s\n", shift, string(decoded))
	}
}

func main() {
	cipherText := "Ugew gnwj zwjw Oslkgf"
	bruteForceCaesarCipher(cipherText)
}
