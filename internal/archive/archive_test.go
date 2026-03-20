package archive

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/fvaiiii/archiver/internal/lz77"
)

func TestFullPipelineRoundTrip(t *testing.T) {
	testCases := []struct {
		name string
		data []byte
	}{
		{"пустой", []byte{}},
		{"один байт", []byte{42}},
		{"короткая строка", []byte("hello")},
		{"повторяющаяся", bytes.Repeat([]byte("ABCD"), 500)},
		{"длинный повтор", bytes.Repeat([]byte("xyz"), 20000)},
		{"реалистичный текст", []byte("Round-trip test (тест кругового пути) — это метод проверки целостности данных, который заключается в преобразовании данных из одного формата в другой.")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tmpFile := filepath.Join(t.TempDir(), "test-"+tc.name+".arc")
			tokens := lz77.Compress(tc.data, lz77.DefaultWindowSize)

			err := WriteArchive(tmpFile, tokens)
			if err != nil {
				t.Fatalf("WriteArchive failed: %v", err)
			}

			readTokens, err := ReadArchive(tmpFile)
			if err != nil {
				t.Fatalf("ReadArchive failed: %v", err)
			}

			decompressed := lz77.Decompress(readTokens)

			if !bytes.Equal(tc.data, decompressed) {
				t.Errorf("Round-trip failed for %q", tc.name)
				t.Errorf("original length: %d, decompressed length: %d", len(tc.data), len(decompressed))
			}

			if len(readTokens) == 0 && len(tc.data) > 0 {
				t.Error("Got zero tokens for non-empty input")
			}
		})
	}
}
