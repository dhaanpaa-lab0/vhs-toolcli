package utils

import (
	"os"
)

func Exists(name string) bool {
	if _, err := os.Stat(name); err == nil {
		return true
	} else {
		return false
	}
}
