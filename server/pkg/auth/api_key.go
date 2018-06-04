package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func UserByApiKey(secretKey string, token string) (int32, string, error) {
	cipherText, _ := base64.URLEncoding.DecodeString(token)

	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return -1, "", err
	}

	if len(cipherText) < aes.BlockSize {
		return -1, "", fmt.Errorf("Invalid length")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(cipherText, cipherText)

	text := string(cipherText)

	arr := strings.Split(text, "@")
	if len(arr) <= 1 {
		return -1, "", fmt.Errorf("Invalid key")
	}

	idPart := arr[0]
	name := strings.Join(arr[1:], "@")

	i, err := strconv.Atoi(idPart)
	if err != nil {
		return -1, "", err
	}

	return int32(i), name, nil
}

func ApiKeyByUser(secretKey string, id int32, name string) (string, error) {
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}

	text := strconv.Itoa(int(id)) + "@" + name
	plainText := []byte(text)
	cipherText := make([]byte, aes.BlockSize+len(plainText))

	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return base64.URLEncoding.EncodeToString(cipherText), nil
}
