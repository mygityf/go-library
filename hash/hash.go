package hash


import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// Sha1 sha1 hash impl
func Sha1(s string) string {
	o := sha1.New()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

// Sha224
func Sha224(s string) string {
	o := sha256.New224()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

// Sha256
func Sha256(s string) string {
	o := sha256.New()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

// Sha384
func Sha384(s string) string {
	o := sha512.New384()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

// Sha512
func Sha512(s string) string {
	o := sha512.New()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

// Md5 hash
func Md5(s string) string {
	o := md5.New()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

