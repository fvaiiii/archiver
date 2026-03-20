package archive

import (
	"encoding/binary"
	"fmt"
	"os"

	"github.com/fvaiiii/archiver/internal/lz77"
)

func WriteArchive(path string, tokens []lz77.Token) error {

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Error opening file: %w", err)
	}

	defer file.Close()

	for _, token := range tokens {
		if token.Length == 0 {
			file.Write([]byte{0, token.NextByte})
		} else {
			file.Write([]byte{1})
			binary.Write(file, binary.LittleEndian, token.Offset)
			binary.Write(file, binary.LittleEndian, token.Length)
		}
	}
	return nil

}
