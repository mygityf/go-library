package hash

import "testing"

func TestAllHash(t *testing.T) {
	args := []struct {
		Raw    string
		Expect string
		Hash   func(string) string
	}{
		{
			Raw:    "123",
			Expect: "202cb962ac59075b964b07152d234b70",
			Hash:   Md5,
		}, {
			Raw:    "123",
			Expect: "40bd001563085fc35165329ea1ff5c5ecbdbbeef",
			Hash:   Sha1,
		}, {
			Raw:    "abc",
			Expect: "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad",
			Hash:   Sha256,
		}, {
			Raw:    "abc",
			Expect: "23097d223405d8228642a477bda255b32aadbce4bda0b3f7e36c9da7",
			Hash:   Sha224,
		}, {
			Raw:    "abc",
			Expect: "cb00753f45a35e8bb5a03d699ac65007272c32ab0eded1631a8b605a43ff5bed8086072ba1e7cc2358baeca134c825a7",
			Hash:   Sha384,
		}, {
			Raw:    "abc",
			Expect: "ddaf35a193617abacc417349ae20413112e6fa4e89a97ea20a9eeee64b55d39a2192992a274fc1a836ba3c23a3feebbd454d4423643ce80e2a9ac94fa54ca49f",
			Hash:   Sha512,
		},
	}
	for _, arg := range args {
		actual := arg.Hash(arg.Raw)
		if actual != arg.Expect {
			t.Errorf("expect:%v, actual:%v", arg.Expect, actual)
		}

	}
}
