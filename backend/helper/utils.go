package helper

import "encoding/base64"

func Base64Encode(str string) string {
	data := []byte(str)
	return base64.StdEncoding.EncodeToString(data)
}
