package utils

import "encoding/base64"

func EncodeBase64Str(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func DecodeBase64Str(str string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
