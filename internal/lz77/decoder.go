package lz77

func Decompress(tokens []Token) []byte {
	var ans []byte

	for _, token := range tokens {
		if token.Length == 0 {
			ans = append(ans, token.NextByte)
		} else {
			if int(token.Offset) > len(ans) {
				panic("invalid offset: past beginning of output")
			}

			start := len(ans) - int(token.Offset)

			for i := 0; i < int(token.Length); i++ {
				if start+i >= len(ans) {
					panic("invalid match length")
				}
				ans = append(ans, ans[start+i])
			}

		}
	}

	return ans
}
