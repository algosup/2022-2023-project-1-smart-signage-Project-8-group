package main

import (
	"encoding/base64"
)

func toBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
