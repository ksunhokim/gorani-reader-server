package view_test

import "testing"

func TestUploadEpub(t *testing.T) {
	token := initDB()
	server, e := initServer(t)
	defer server.Close()

	e.POST("/upload").
		WithHeader("Authorization", token).
		WithMultipart().
		WithFile("epub", "./test.epub").
		Expect().
		Status(200)
}
