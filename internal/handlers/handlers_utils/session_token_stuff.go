package handlers_utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"log"
)

// CipherString creates an encrypted string from provided data(string) and secret(string)
func CipherString(inputString string, secretKey string) (string, error) {
	// cipher creation
	generatedCipher, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}
	// wrap generatedCipher with 'GCM' to authenticate (protect) the data
	gcm, err := cipher.NewGCM(generatedCipher)
	if err != nil {
		return "", err
	}
	// NONCE creation
	// getting a size of "used once number"
	nonce_size := gcm.NonceSize()
	// bytes array creation to fill in
	nonce := make([]byte, nonce_size)
	// generate (read from rand.Reader) random bytes for nonce
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	// finally encrypt it!
	cipherTextBytes := gcm.Seal(nonce, nonce, []byte(inputString), nil)

	// convert to string with base64 will be successfull
	// even if some bytes in cipherText are incorrect
	cipherText := base64.StdEncoding.EncodeToString(cipherTextBytes)
	log.Printf("Finished session token encryption. cipherText = %s", cipherText)
	return cipherText, nil
}

func DecryptCipherString(encryptedData string, secretKey string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}

	c, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
