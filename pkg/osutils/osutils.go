package osutils

import (
	"os"
	"strings"
)

// GetEnvWithDefault - get environment variable; if no environment variable then
// return defaultStr.
func GetEnvWithDefault(key string, defaultStr string) string {
	val := strings.TrimSpace(os.Getenv(key))
	if val == "" {
		val = defaultStr
	}
	return val
}
