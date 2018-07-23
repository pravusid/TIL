package hangul

var (
	start = rune(44032) // "가"의 유니코드
	end   = rune(55204) // "힣" 다음글자의 유니코드
)

// HasConsonantSuffix는 받침이 있는 한글 글자인 경우 true 반환
func HasConsonantSuffix(s string) bool {
	numEnds := 28
	result := false
	for _, r := range s {
		if start <= r && r < end {
			index := int(r - start)
			result = index%numEnds != 0
		}
	}
	return result
}
