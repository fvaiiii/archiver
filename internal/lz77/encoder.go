package lz77

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

func Compress(data []byte, windowSize int) []Token {
	var tokens []Token
	n := len(data)
	pos := 0

	const maxMathLen = 65535

	for pos < n {
		matchOffset := 0
		matchLength := 0

		windowStart := max(0, pos-windowSize)

		for i := windowStart; i < pos; i++ {
			if matchLength >= maxMathLen {
				break
			}

			lenMatch := 0
			for pos+lenMatch < n && i+lenMatch < pos && data[i+lenMatch] == data[pos+lenMatch] {
				lenMatch++
				if lenMatch+pos-i >= matchLength+1 {
					break
				}
			}

			if lenMatch > matchLength {
				matchLength = lenMatch
				matchOffset = pos - i
			}
		}

		if matchLength >= 3 {
			tokens = append(tokens, Token{
				Offset:   uint16(matchOffset),
				Length:   uint16(min(matchLength, maxMathLen)),
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
