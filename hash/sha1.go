package hash

import (
	"crypto/sha1"
	"encoding/hex"
)

// Sha1 sha1 hash impl
func Sha1(s string) string {
	o := sha1.New()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}
