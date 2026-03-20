package archive

import (
	"encoding/binary"
	"os"

	"github.com/fvaiiii/archiver/internal/lz77"
)

func ReadArchive(path string) ([]lz77.Token, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var tokens []lz77.Token
	for i := 0; i < len(data); {
		flag := data[i]
		i++
		if flag == 0 {
			tokens = append(tokens, lz77.Token{NextByte: data[i]})
			i++
		} else {
			offset := binary.LittleEndian.Uint16(data[i : i+2])
			length := binary.LittleEndian.Uint16(data[i+2 : i+4])
			tokens = append(tokens, lz77.Token{Offset: offset, Length: length})
			i += 4
		}
	}
	return tokens, nil
}
