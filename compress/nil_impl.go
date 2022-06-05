package compress

type nilImpl struct {
}

// New
func NewNilImpl() *nilImpl {
	return &nilImpl{}
}

// encode
func (g *nilImpl) Encode(src []byte) ([]byte, error) {
	return src, nil
}

// decode
func (g *nilImpl) Decode(src []byte) ([]byte, error) {
	return src, nil
}
