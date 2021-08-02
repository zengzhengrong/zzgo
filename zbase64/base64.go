package zbase64

import (
	b64 "encoding/base64"
)

// Encode 编码
func Encode(s string) string {
	return b64.StdEncoding.EncodeToString([]byte(s))
}

// Decode 解码
func Decode(s string) (string, error) {
	ds, err := b64.StdEncoding.DecodeString(s)
	return string(ds), err
}
