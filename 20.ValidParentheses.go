package main

import "fmt"

func isValid(s string) bool {
	result := false

	if len(s)%2 != 0 {
		return result
	}

	return result
}

func main() {
	fmt.Println(isValid("()"))     // Output: true
	fmt.Println(isValid("()[]{}")) // Output: true
	fmt.Println(isValid("(]"))     // Output: false
	fmt.Println(isValid("([)]"))   // Output: false
	fmt.Println(isValid("{[]}"))   // Output: true
}
