package archive

import (
	"encoding/binary"
	"fmt"
	"os"

	"github.com/fvaiiii/archiver/internal/lz77"
)

func WriteArchive(path string, tokens []lz77.Token) error {

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("Ошибка открытия файла: %w", err)
	}

	defer file.Close()

	_, err = file.Write([]byte("ARCV"))
	if err != nil {
		return err
	}

	err = binary.Write(file, binary.LittleEndian, uint32(len(tokens)))
	if err != nil {
		return fmt.Errorf("Error writing token count: %w", err)
	}

	for _, token := range tokens {
		err = binary.Write(file, binary.LittleEndian, token.Offset)
		if err != nil {
			return err
		}
		err = binary.Write(file, binary.LittleEndian, token.Length)
		if err != nil {
			return err
		}
		err = binary.Write(file, binary.LittleEndian, token.NextByte)
		if err != nil {
			return err
		}
	}
	return nil

}
