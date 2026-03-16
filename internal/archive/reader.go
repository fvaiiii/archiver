package archive

import (
	"encoding/binary"
	"fmt"
	"os"

	"github.com/fvaiiii/archiver/internal/lz77"
)

func ReadArchive(path string) ([]lz77.Token, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Error opening archive %w", err)
	}
	defer file.Close()

	header := make([]byte, 4)
	_, err = file.Read(header)
	if err != nil {
		return nil, err
	}

	if string(header) != "ARCV" {
		return nil, fmt.Errorf("Invalid archive format")
	}

	var tokenCount uint32
	err = binary.Read(file, binary.LittleEndian, &tokenCount)
	if err != nil {
		return nil, err
	}

	tokens := make([]lz77.Token, 0, tokenCount)
	for i := uint32(0); i < tokenCount; i++ {
		var token lz77.Token
		err = binary.Read(file, binary.LittleEndian, &token.Offset)
		if err != nil {
			return nil, err
		}

		err = binary.Read(file, binary.LittleEndian, &token.Length)
		if err != nil {
			return nil, err
		}

		err = binary.Read(file, binary.LittleEndian, &token.NextByte)
		if err != nil {
			return nil, err
		}

		tokens = append(tokens, token)
	}
	return tokens, nil
}
