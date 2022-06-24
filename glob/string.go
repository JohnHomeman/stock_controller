package glob

import (
	"encoding/base64"
	"strconv"
)

func DecodeBase64(s string) string {
	dec, _ := base64.StdEncoding.DecodeString(s)
	return string(dec)
}

func IsInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
