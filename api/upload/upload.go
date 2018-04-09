package upload

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"strings"

	"github.com/sunho/engbreaker/pkg/dbs"
)

func ToBytes(header *multipart.FileHeader) ([]byte, error) {
	if !strings.HasSuffix(header.Filename, ".epub") {
		return []byte{}, fmt.Errorf("File not supported")
	}
	file, err := header.Open()
	if err != nil {
		return []byte{}, err
	}
	bytes, err := ioutil.ReadAll(file)
	return bytes, err
}

func UploadToRedis(bytes []byte) (string, error) {
	nonce := dbs.GenerateRedisNonce("upload:")
	err := dbs.RDB.Set(nonce, bytes, 0)
	if err != nil {
		return "", err.Err()
	}
	return nonce, nil
}
