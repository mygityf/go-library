package httplib

import (
	"compress/gzip"
	"net"
)

type gzipConn struct {
	net.Conn
	*gzip.Writer
}

// New
func NewGzipConn(conn net.Conn) *gzipConn {
	return &gzipConn{
		Conn:   conn,
		Writer: gzip.NewWriter(conn),
	}
}

// Write
func (gc gzipConn) Write(b []byte) (int, error) {
	n, err := gc.Writer.Write(b)
	if err != nil {
		return n, err
	}

	return n, gc.Writer.Flush()
}

// Close
func (gc gzipConn) Close() error {
	err := gc.Writer.Close()
	if err != nil {
		return err
	}

	return gc.Conn.Close()
}
