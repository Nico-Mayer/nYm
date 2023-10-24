package utils

import (
	"crypto/md5"
	"encoding/base64"
	"net/url"
)

func IsValidURL(inputURL string) bool {
	_, err := url.ParseRequestURI(inputURL)
	return err == nil
}

func EncryptURL(url string) string {
	hash := md5.Sum([]byte(url))
	return base64.URLEncoding.EncodeToString(hash[:])[:6]
}
