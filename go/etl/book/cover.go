package book

import (
	"errors"

	"github.com/google/uuid"
	minio "github.com/minio/minio-go"
)

var (
	ErrNoCover = errors.New("book: the epub doesn't have cover")
)

func (b *Book) UploadCover(m *minio.Client) error {
	if b.Cover.Reader == nil {
		return ErrNoCover
	}

	name := uuid.New().String() + "." + b.Cover.Ext
	_, err := m.PutObject("picture", name, b.Cover.Reader, -1, minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	b.Cover.Object = name

	return nil
}
