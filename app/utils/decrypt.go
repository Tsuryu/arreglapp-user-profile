package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
)

// Decrypt : decript a string
func Decrypt(value string) (string, error) {
	data := []byte(value)
	key, err := generateKey()
	if err != nil {
		log.Fatal(err)
	}

	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return string(""), err
	}
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return string(""), err
	}
	nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return string(""), err
	}
	return string(plaintext), nil
}
