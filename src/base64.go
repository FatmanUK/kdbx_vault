package main

import (
	b64 "encoding/base64"
)

func Base64Decode(b64str string) string {
	plain, err := b64.StdEncoding.DecodeString(b64str)
	if err != nil {
		panic(err)
	}
	return string(plain)
}
