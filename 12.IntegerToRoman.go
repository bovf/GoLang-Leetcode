package golangleetcode

func intToRoman(num int) string {
	carryOver := num
	result := ""
	numerals := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for carryOver > 0 {
		if carryOver >= 1000 {
			result += "M"
			carryOver -= numerals["M"]
			continue
		}
		if carryOver < 1000 && carryOver >= 900 {
			result += "CM"
			carryOver -= 900
			continue
		}
		if carryOver >= 500 {
			result += "D"
			carryOver -= numerals["D"]
			continue
		}
		if carryOver < 500 && carryOver >= 400 {
			result += "CD"
			carryOver -= 400
			continue
		}
		if carryOver >= 100 {
			result += "C"
			carryOver -= numerals["C"]
			continue
		}
		if carryOver < 100 && carryOver >= 90 {
			result += "XC"
			carryOver -= 90
			continue
		}
		if carryOver >= 50 {
			result += "L"
			carryOver -= numerals["L"]
			continue
		}
		if carryOver < 50 && carryOver >= 40 {
			result += "XL"
			carryOver -= 40
			continue
		}
		if carryOver >= 10 {
			result += "X"
			carryOver -= numerals["X"]
			continue
		}
		if carryOver < 10 && carryOver >= 9 {
			result += "IX"
			carryOver -= 9
			continue
		}
		if carryOver >= 5 {
			result += "V"
			carryOver -= numerals["V"]
			continue
		}
		if carryOver < 5 && carryOver >= 4 {
			result += "IV"
			carryOver -= 4
			continue
		}
		if carryOver >= 1 {
			result += "I"
			carryOver -= numerals["I"]
			continue
		}
	}

	return result
}
