package utils

import (
	"net/http"
	"os"
	"path/filepath"
)

// CrossPlatformFilepath ensures that paths are compatible with the OS, e.g.:
// * a\b\c will be a/b/c on linux
// * d/e/f will be d\e\f on Windows
func CrossPlatformFilepath(p string) string {
	return filepath.FromSlash(p)
}

// ToLinuxPath transforms back slashes to forward slashes
func ToLinuxPath(p string) string {
	return filepath.ToSlash(p)
}

// URLExists checks whether an URL exists
func URLExists(url string) bool {
	_, err := http.Get(url)
	if err != nil {
		return false
	}
	return true
}

// FileExists checks whether a file exists
func FileExists(f string) bool {
	_, err := os.Stat(f)
	if err != nil {
		return false
	}
	return true
}
