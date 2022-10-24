package utils

import (
	"os"
	"os/user"
	"strconv"
	"strings"
)

func StringFromEnv(envName string, defaultStrVal string) string {
	strVal := strings.TrimSpace(os.Getenv(envName))
	if strVal == "" {
		strVal = defaultStrVal
	}
	return strVal
}

func IntFromEnv(envName string, defaultIntVal int) int {
	strVal := strings.TrimSpace(os.Getenv(envName))
	if strVal == "" {
		return defaultIntVal
	}
	intVal, err := strconv.Atoi(strVal)
	if err != nil {
		return defaultIntVal
	}
	return intVal
}

func Username(defaultUsername string) string {
	currentUser, err := user.Current()
	if err != nil {
		return defaultUsername
	}
	return currentUser.Username
}
