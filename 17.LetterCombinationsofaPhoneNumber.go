package main

func backTrack(digitMap map[string][]string, currentString string, result *[]string, i int, digits string, arraySize int, arrayIndex *int) {
	if len(currentString) == len(digits) {
		(*result)[*arrayIndex] = currentString
		*arrayIndex++
		return
	}

	chars := digitMap[string(digits[i])]

	for _, char := range chars {
		backTrack(digitMap, currentString+char, result, i+1, digits, arraySize, arrayIndex)
	}
}

func getArraySize(digitMap map[string][]string, digits string) int {
	size := 1

	for _, digit := range digits {
		size *= len(digitMap[string(digit)])
	}

	return size
}

func letterCombinations(digits string) []string {
	numPad := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	if len(digits) == 0 {
		return []string{}
	}
	arrayIndex := 0
	arraySize := getArraySize(numPad, digits)
	result := make([]string, arraySize)
	backTrack(numPad, "", &result, 0, digits, arraySize, &arrayIndex)
	return result
}
