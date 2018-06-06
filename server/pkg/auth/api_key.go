package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"strconv"
	"strings"
)

var (
	ErrInvalidLength = errors.New("auth: Invalid length")
	ErrInvalidKey    = errors.New("auth: Invalid key")
)

// the name from a valid api key is equal to the actual name in db
func UserByApiKey(secretKey string, token string) (id int32, name string, err error) {
	cipherText, _ := base64.URLEncoding.DecodeString(token)

	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return
	}

	if len(cipherText) < aes.BlockSize {
		err = ErrInvalidLength
		return
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(cipherText, cipherText)

	text := string(cipherText)

	arr := strings.Split(text, "@")
	if len(arr) <= 1 {
		err = ErrInvalidKey
		return
	}

	idPart := arr[0]
	name = strings.Join(arr[1:], "@")

	i, err := strconv.Atoi(idPart)
	if err != nil {
		return
	}

	id = int32(i)
	return
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
