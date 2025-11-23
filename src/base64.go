package main

import (
	"strings"
	b64 "encoding/base64"
)

func mustBase64Decode(b64str string) string {
	plain, err := b64.StdEncoding.DecodeString(b64str)
	if err != nil {
		panic(err)
	}
	return string(plain)
}

func base64Decode(b64str string) (string, error) {
	decodedBytes, err := b64.StdEncoding.DecodeString(b64str)
	decodedString := strings.TrimSpace(string(decodedBytes))
	return decodedString, err
}
