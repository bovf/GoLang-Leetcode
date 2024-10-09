package main

func isValid(s string) bool {
	var result bool

	if len(s)%2 != 0 {
		result = false
		return result
	}

	stack := []rune{}
	openBrackets := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, r := range s {
		if closer, ok := openBrackets[r]; ok {
			stack = append(stack, closer)
			continue
		}

		stackLength := len(stack)
		if stackLength == 0 || r != stack[stackLength-1] {
			result = false
			return result
		}

		stack = stack[:stackLength-1]
	}
	result = len(stack) == 0
	return result
}
