package auth

import "io"

type User struct {
	Username string
	Avator   string
	Id       string
}

type Service interface {
	Name() string
	BaseUrl() string
	UserEndpoint() string
	ReaderToUser(io.Reader) (User, error)
}
