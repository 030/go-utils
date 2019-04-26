package utils

import "testing"

func TestURLExists(t *testing.T) {
	urlHashMap := map[string]bool{
		"http://google.com":          true,
		"http://google444444444.com": false,
	}

	for url, expectedExistence := range urlHashMap {
		actualExistence := URLExists(url)
		if expectedExistence != actualExistence {
			t.Errorf("Expected that URL %v existence would be %v, but was %v", url, expectedExistence, actualExistence)
		}
	}
}
