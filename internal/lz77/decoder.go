package lz77

func Decompress(tokens []Token) []byte {
	var ans []byte
	for _, token := range tokens {
		if token.Length > 0 {
			start := len(ans) - token.Offset

			for i := 0; i < token.Length; i++ {
				ans = append(ans, ans[start+i])
			}
		}
		if token.NextByte != 0 {
			ans = append(ans, token.NextByte)
		}

	}
	return ans
}
