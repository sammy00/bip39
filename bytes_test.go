package bip39

import (
	"bytes"
	"testing"
)

func Test_paddedPrepend(t *testing.T) {
	testCases := []struct {
		dst, src []byte
		size     int
		expect   []byte
	}{
		{nil, []byte{1, 2, 3}, 4, []byte{0, 1, 2, 3}},
		{[]byte{5}, []byte{1, 2, 3}, 4, []byte{5, 0, 1, 2, 3}},
		{[]byte{5}, []byte{1, 2, 3}, 3, []byte{5, 1, 2, 3}},
		{[]byte{5}, []byte{1, 2, 3}, 2, []byte{5, 1, 2, 3}},
	}

	for i, c := range testCases {
		if got := paddedPrepend(c.dst, c.src, c.size); !bytes.Equal(got, c.expect) {
			t.Fatalf("#%d failed: got %v, expect %v", i, got, c.expect)
		}
	}
}
