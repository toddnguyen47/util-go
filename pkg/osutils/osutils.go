package osutils

import (
	"os"
	"strconv"
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

func GetEnvWithDefaultInt(key string, defaultInt int) int {
	val := strings.TrimSpace(os.Getenv(key))
	valInt, err := strconv.Atoi(val)
	if err != nil {
		return defaultInt
	}
	return valInt
}

func GetEnvWithDefaultFloat64(key string, defaultFloat float64) float64 {
	val := strings.TrimSpace(os.Getenv(key))
	valInt, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return defaultFloat
	}
	return valInt
}

func GetEnvWithDefaultBool(key string, defaultBool bool) bool {
	val := strings.TrimSpace(os.Getenv(key))
	valBool, err := strconv.ParseBool(val)
	if err != nil {
		return defaultBool
	}
	return valBool
}

func RemoveIfExists(path string) error {
	if _, err := os.Stat(path); err == nil {
		return os.Remove(path)
	}
	return nil
}
