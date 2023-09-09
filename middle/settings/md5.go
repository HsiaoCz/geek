package settings

import (
	"crypto/md5"
	"encoding/hex"
)

var secret = "xiaofanyi"

func SetMd5Password(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
