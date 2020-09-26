package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"log"
)

// // Encrypt : encrypts a string
// func encrypt(value string) (string, error) {
// 	encryptedBytes, err := bcrypt.GenerateFromPassword([]byte(value), 10)
// 	if err != nil {
// 		return string(""), err
// 	}
// 	return string(encryptedBytes), nil
// }

// Encrypt : encrypts a string
func Encrypt(value string) (string, error) {
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

	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return string(""), err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	return hex.EncodeToString(ciphertext), nil
}
