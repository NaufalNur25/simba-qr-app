package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"
)

func EncryptAES(plainText string) (string, error) {
	var secretKey = []byte(os.Getenv("AES_SECRET_KEY"))

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	plainBytes := []byte(plainText)
	cipherText := make([]byte, aes.BlockSize+len(plainBytes))

	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainBytes)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}
