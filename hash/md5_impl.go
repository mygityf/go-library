package hash

import (
	"crypto/md5"
	"encoding/hex"
)

type md5Impl struct {
	key []byte
}

// New
func NewMd5() *md5Impl {
	return &md5Impl{}
}

// New with key
func NewMd5WithKey(key []byte) *md5Impl {
	return &md5Impl{key: key}
}

// Hash
func (m *md5Impl) HashEncodeHex(src []byte) string {
	hashImpl := md5.New()
	if m.key != nil {
		hashImpl.Write(m.key)
	}
	hashImpl.Write(src)
	res := hashImpl.Sum(nil)
	return hex.EncodeToString(res)
}

// Hash
func (m *md5Impl) HashRaw(src []byte) []byte {
	hashImpl := md5.New()
	if m.key != nil {
		hashImpl.Write(m.key)
	}
	hashImpl.Write(src)
	return hashImpl.Sum(nil)
}
