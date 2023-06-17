package util

import "encoding/base64"

func ToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
