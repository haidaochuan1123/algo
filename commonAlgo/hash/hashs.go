package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

// HashBySHA1 通过SHA-1对内容加密
func HashBySHA1(b []byte) []byte {
	h := sha1.New()
	h.Write(b)
	res := h.Sum(nil)
	return res
}

// HashByMD5 通过MD5对内容加密
func HashByMD5(b []byte) string {
	h := md5.New()
	h.Write(b)
	res := h.Sum(nil)
	result := hex.EncodeToString(res)
	return result
}

// HashByMD5Simple 通过MD5对内容加密
func HashByMD5Simple(b []byte) string {
	res := md5.Sum(b)
	result := hex.EncodeToString(res[:])
	return result
}
