package util

import "os"

func GetEnvOr(name, value string) string {
	v := os.Getenv(name)
	if v == "" {
		return value
	}
	return v
}
