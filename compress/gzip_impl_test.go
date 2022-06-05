package compress

import "testing"

func TestGzipImpl_Encode(t *testing.T) {
	impl := NewGzipImpl()
	encodedBytes, err := impl.Encode([]byte("11223344"))
	if err != nil {
		t.Error(err)
	}
	t.Logf("%x", encodedBytes)
	decodedBytes, err := impl.Decode(encodedBytes)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(decodedBytes))
}
