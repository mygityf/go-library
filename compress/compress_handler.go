package compress

type ComoressHandler interface {
	// encode
	Encode([]byte) ([]byte, error)
	// decode
	Decode([]byte) ([]byte, error)
}