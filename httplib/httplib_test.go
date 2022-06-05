package httplib

import (
	"testing"
)

func TestGet(t *testing.T) {
	req := Get("https://www.qq.com")
	b, err := req.Bytes()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(b)
}
