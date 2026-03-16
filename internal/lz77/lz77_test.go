package lz77

import (
	"bytes"
	"testing"
)

type TestCase struct {
	name string
	text []byte
}

func TestCompressAndDecompress(t *testing.T) {
	cases := []TestCase{
		{"test1", []byte{}},
		{"test2", []byte{65}},
		{"test3", []byte("abcdef")},
		{"test4", []byte("aaaaaa")},
		{"test5", []byte("abcabcabc")},
		{"test6", []byte("aacaacabcabaaac")},
		{"test7", []byte("hellohello world")},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			tokens := Compress(c.text, 64)
			result := Decompress(tokens)

			if !bytes.Equal(c.text, result) {
				t.Fatalf("original and result are not equal")
			}
		})
	}
}
