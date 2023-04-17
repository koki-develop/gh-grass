package cmd

import "os"

func getEnvOr(name, value string) string {
	v := os.Getenv(name)
	if v == "" {
		return value
	}
	return v
}
