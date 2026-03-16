package lz77

func Compress(data []byte, windowSize int) []Token {
	var tokens []Token

	pos := 0
	n := len(data)

	for pos < n {
		matchOffset := 0
		matchLength := 0

		windowStart := pos - windowSize
		if windowStart < 0 {
			windowStart = 0
		}

		for j := windowStart; j < pos; j++ {
			length := 0

			for j+length < n && data[j+length] == data[pos+length] {
				length++
				if j+length >= pos {
					break
				}
			}

			if length > matchLength {
				matchLength = length
				matchOffset = pos - j
			}
		}

		if matchLength > 0 {
			var nextByte byte
			if pos+matchLength < n {
				nextByte = data[pos+matchLength]
			}
			tokens = append(tokens, Token{
				Offset:   uint16(matchOffset),
				Length:   uint16(matchLength),
				NextByte: nextByte,
			})

			pos += matchLength + 1
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
