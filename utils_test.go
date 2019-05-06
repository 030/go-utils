package utils

import "testing"

func TestURLExists(t *testing.T) {
	urlHashMap := map[string]bool{
		"http://releasesoftwaremoreoften.com":           true,
		"http://releasesoftwaremoreoften4444444444.com": false,
	}

	for url, expectedExistence := range urlHashMap {
		actualExistence := URLExists(url)
		if expectedExistence != actualExistence {
			t.Errorf("Expected that URL %v existence would be %v, but was %v", url, expectedExistence, actualExistence)
		}
	}
}

func TestFileExists(t *testing.T) {
	fileHashMap := map[string]bool{
		"go.mod":               true,
		"go.mod2":              false,
		"/usr/local/something": false,
	}

	for file, expectedExistence := range fileHashMap {
		actualExistence := FileExists(file)
		if expectedExistence != actualExistence {
			t.Errorf("Expected that file %v existence would be %v, but was %v", file, expectedExistence, actualExistence)
		}
	}
}
