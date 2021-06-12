package cryptox

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 md5 字符串摘要
func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}
