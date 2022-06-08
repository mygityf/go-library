package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

// Sha1 sha1 hash impl
func Sha1(s string) string {
	o := sha1.New()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

// Md5 hash
func Md5(s string) string {
	o := md5.New()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

