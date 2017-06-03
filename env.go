package toolkit

import "os"

func Env(key, defaultValue string) string {
	v := os.Getenv(key)
	if len(v) == 0 {
		return defaultValue
	}
	return v
}
