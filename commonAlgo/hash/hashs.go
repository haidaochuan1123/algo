package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

// SHA1Hash 通过SHA-1对内容加密
func SHA1Hash(b []byte) []byte {
	h := sha1.New()
	h.Write(b)
	res := h.Sum(nil)
	return res
}

// MD5Hash 通过MD5对内容加密
func MD5Hash(b []byte) string {
	h := md5.New()
	h.Write(b)
	res := h.Sum(nil)
	result := hex.EncodeToString(res)
	return result
}

// MD5SimpleHash 通过MD5对内容加密
func MD5SimpleHash(b []byte) string {
	res := md5.Sum(b)
	result := hex.EncodeToString(res[:])
	return result
}
