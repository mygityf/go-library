package compress

import (
	"bytes"
	"compress/gzip"
	"io"
)

type gzipImpl struct {
}

// New
func NewGzipImpl() *gzipImpl {
	return &gzipImpl{}
}

// encode
func (g *gzipImpl) Encode(src []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	if _, err := zw.Write(src); err != nil {
		return nil, err
	}
	zw.Close()
	return buf.Bytes(), nil
}

// decode
func (g *gzipImpl) Decode(src []byte) ([]byte, error) {
	var buf bytes.Buffer
	buf.Write(src)
	zr, err := gzip.NewReader(&buf)
	if err != nil {
		return nil, err
	}
	defer zr.Close()
	var destBuf bytes.Buffer
	if _, err = io.Copy(&destBuf, zr); err != nil {
		return nil, err
	}
	return destBuf.Bytes(), nil
}
