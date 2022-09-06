package outil

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Encode(in string) string {
	md5Value := md5.New()
	md5Value.Write([]byte(in))
	return hex.EncodeToString(md5Value.Sum([]byte{}))
}
