package lz77

func Decompress(tokens []Token) []byte {
	var ans []byte
	for _, token := range tokens {
		if token.Length > 0 {
			start := len(ans) - int(token.Offset)

			for i := 0; i < int(token.Length); i++ {
				ans = append(ans, ans[start+i])
			}
		}
		ans = append(ans, token.NextByte)

	}
	return ans
}
