package hash

import "testing"

func TestAllHmac(t *testing.T) {
	args := []struct {
		Raw    string
		Key    []byte
		Expect string
		Hash   func([]byte, string) string
	}{
		{
			Raw:    "The quick brown fox jumps over the lazy dog",
			Key:    []byte("key"),
			Expect: "80070713463e7749b90c2dc24911e275",
			Hash:   Md5Hmac,
		}, {
			Raw:    "wangyaofu try sha1 hash.",
			Key:    []byte("key"),
			Expect: "9b3b3eafabef9a054b28cc580be523793dc7f461",
			Hash:   Sha1Hmac,
		}, {
			Raw:    "wangyaofu try sha256 hash.",
			Key:    []byte("key"),
			Expect: "dc1099cd35452f3cf1f6c11d3884bc4c9b637523c4e74060e1332eac01e7ad8e",
			Hash:   Sha256Hmac,
		}, {
			Raw:    "wangyaofu try sha224 hash.",
			Key:    []byte("key"),
			Expect: "4d5a9a6ed549145dfbc30b0752ae3e05e06ccdd5695ec980f5a5739e",
			Hash:   Sha224Hmac,
		}, {
			Raw:    "abc",
			Key:    []byte("key"),
			Expect: "30ddb9c8f347cffbfb44e519d814f074cf4047a55d6f563324f1c6a33920e5edfb2a34bac60bdc96cd33a95623d7d638",
			Hash:   Sha384Hmac,
		}, {
			Raw:    "abc",
			Key:    []byte("key"),
			Expect: "3926a207c8c42b0c41792cbd3e1a1aaaf5f7a25704f62dfc939c4987dd7ce060009c5bb1c2447355b3216f10b537e9afa7b64a4e5391b0d631172d07939e087a",
			Hash:   Sha512Hmac,
		},
	}
	for _, arg := range args {
		actual := arg.Hash(arg.Key, arg.Raw)
		if actual != arg.Expect {
			t.Errorf("expect:%v, actual:%v", arg.Expect, actual)
		}

	}
}
