package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat:/ama.me" {
		return false
	}
	return true
}

// func TestChe
