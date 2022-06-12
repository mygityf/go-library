package hash

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// Sha1Hmac sha1 hash impl
func Sha1Hmac(key []byte, s string) string {
	h := hmac.New(sha1.New, key)
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha224 hash
func Sha224Hmac(key []byte, s string) string {
	h := hmac.New(sha256.New224, key)
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha256 hash
func Sha256Hmac(key []byte, s string) string {
	h := hmac.New(sha256.New, key)
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha384 hash
func Sha384Hmac(key []byte, s string) string {
	h := hmac.New(sha512.New384, key)
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha512 hash
func Sha512Hmac(key []byte, s string) string {
	h := hmac.New(sha512.New, key)
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Md5Hmac hash
func Md5Hmac(key []byte, s string) string {
	h := hmac.New(md5.New, key)
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
