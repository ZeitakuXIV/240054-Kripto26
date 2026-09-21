package main

import "fmt"

var chars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func encryptVigenere(text, key string) string {
	result := ""
	for i := 0; i < len(text); i++ {
		textPos := int(text[i] - 'A')
		keyPos := int(key[i%len(key)] - 'A')

		shift := (textPos + keyPos) % 26
		encryptedChar := chars[shift]
		result += string(encryptedChar)
	}
	return result
}

func decryptVigenere(text, key string) string {
	result := ""
	for i := 0; i < len(text); i++ {
		textPos := int(text[i] - 'A')
		keyPos := int(key[i%len(key)] - 'A')

		shift := (textPos - keyPos + 26) % 26
		decryptedChar := chars[shift]
		result += string(decryptedChar)
	}
	return result
}

func main() {
	text := "ASPRAKGANTENG"
	key := "MNURRIZALZIDM"

	encV := encryptVigenere(text, key)
	fmt.Println("Vigenere Encrypted:", encV)
	fmt.Println("Vigenere Key:", key)
	fmt.Println("Vigenere Decrypted:", decryptVigenere(encV, key))
}
