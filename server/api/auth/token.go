package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"strconv"
	"strings"
)

func IdFromToken(key []byte, token string, name string) int {
	cipherText, _ := base64.URLEncoding.DecodeString(token)

	block, err := aes.NewCipher(key)
	if err != nil {
		return -1
	}

	if len(cipherText) < aes.BlockSize {
		return -1
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(cipherText, cipherText)

	text := string(cipherText)

	arr := strings.Split(text, "@")
	if len(arr) <= 1 {
		return -1
	}

	idPart := arr[0]
	namePart := strings.Join(arr[1:], "@")
	if namePart != name {
		return -1
	}

	i, err := strconv.Atoi(idPart)
	if err != nil {
		return -1
	}

	return i
}

func TokeFromId(key []byte, id int, name string) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		return err.Error()
	}

	text := strconv.Itoa(id) + "@" + name
	plainText := []byte(text)
	cipherText := make([]byte, aes.BlockSize+len(plainText))

	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err.Error()
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return base64.URLEncoding.EncodeToString(cipherText)
}
