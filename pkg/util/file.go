package util

import "os"

func FileExist(url string) bool {
	_, err := os.Stat(url)
	return !os.IsNotExist(err)
}
