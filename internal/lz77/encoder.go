package lz77

const (
	MaxMatchLen       = 65535
	MinMatchLen       = 3
	DefaultWindowSize = 32768
)

func Compress(data []byte, windowSize int) []Token {
	var tokens []Token
	n := len(data)
	pos := 0

	for pos < n {
		matchOffset := 0
		matchLength := 0

		windowStart := max(0, pos-windowSize)

		for i := windowStart; i < pos; i++ {

			lenMatch := 0
			for pos+lenMatch < n && data[i+lenMatch] == data[pos+lenMatch] {
				lenMatch++
				if lenMatch >= MaxMatchLen {
					break
				}
			}

			if lenMatch > matchLength {
				matchLength = lenMatch
				matchOffset = pos - i
			}
		}

		if matchLength >= MinMatchLen {
			tokens = append(tokens, Token{
				Offset:   uint16(matchOffset),
				Length:   uint16(min(matchLength, MaxMatchLen)),
				NextByte: 0,
			})
			pos += matchLength
		} else {

			tokens = append(tokens, Token{
				Offset:   0,
				Length:   0,
				NextByte: data[pos],
			})
			pos++
		}
	}

	return tokens
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
