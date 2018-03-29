package view

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sunho/engbreaker/pkg/upload"
)

func Upload(c *gin.Context) {
	header, err := c.FormFile("epub")
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	bytes, err := upload.ToBytes(header)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	file, err := upload.UploadToRedis(bytes)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	fmt.Println(file)
	c.Status(200)
}
