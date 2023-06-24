package utils

import (
	"fmt"
	"os"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func RemoveDomainError(url string) bool {
	if url == os.Getenv("DOMAIN") {
		return false
	}
	fmt.Println(url, os.Getenv("DOMAIN"))
	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]

	return newURL != os.Getenv("DOMAIN")
}

func EnforceHTTP(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}

func Base62Encode(number uint64) string {
	length := len(alphabet)
	fmt.Println("length of alphabet", length)
	var encodedBuild strings.Builder
	fmt.Println("encodedBuild", encodedBuild)

	encodedBuild.Grow(10)
	for ; number > 0; number = number / uint64(length) {
		encodedBuild.WriteByte(alphabet[(number % uint64(length))])
	}

	return encodedBuild.String()
}
