package config

import "os"

func GetString(name string) string {
	return os.Getenv(name)
}
