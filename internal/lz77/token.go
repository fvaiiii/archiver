package lz77

type Token struct {
	Offset   int
	Length   int
	NextByte byte
}
