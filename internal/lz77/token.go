package lz77

type Token struct {
	Offset   uint16
	Length   uint16
	NextByte byte
}
