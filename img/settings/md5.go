package settings

import (
	"crypto/md5"
	"encoding/hex"
)

var sercet = []byte("xiaofanyi")

// GenMD5 对用户的密码进行加密

func GenMD5(password string) string {
  h:=md5.New()
  h.Write(sercet)
  return hex.EncodeToString(h.Sum([]byte(password)))
}
