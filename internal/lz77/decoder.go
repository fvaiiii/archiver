package lz77

func Decompress(tokens []Token) []byte {
	ans := make([]byte, 0, len(tokens)*2)

	for _, token := range tokens {
		if token.Length == 0 {
			ans = append(ans, token.NextByte)
		} else {
			start := len(ans) - int(token.Offset)
			if start < 0 {
				panic("invalid offset: past beginning of output")
			}

			for i := 0; i < int(token.Length); i++ {
				ans = append(ans, ans[start+i])
			}

		}
	}
	return ans
}
