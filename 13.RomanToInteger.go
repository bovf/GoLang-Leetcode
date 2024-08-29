package golangleetcode

func romanToInt(s string) int {
	result := 0
	numerals := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i := 0; i < len(s); i++ {
		if i+1 <= len(s)-1 {
			first := string(s[i])
			second := string(s[i+1])
			if first < second {
				result += numerals[second] - numerals[first]
				i++
			} else {
				result += numerals[first]
			}
		} else {
			first := string(s[i])
			result += numerals[first]
		}
	}

	return result
}
