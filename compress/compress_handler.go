package compress

type CompressHandler interface {
	// encode
	Encode([]byte) ([]byte, error)
	// decode
	Decode([]byte) ([]byte, error)
}