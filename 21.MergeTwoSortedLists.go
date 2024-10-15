package main

import "strings"

func generateParenthesis(n int) []string {
	combination := []string{}
	stack := []string{}
	if n == 0 {
		return combination
	}
	backtrack(0, 0, n, &combination, &stack)
	return combination
}

func backtrack(open int, closed int, n int, combination *[]string, stack *[]string) {
	if open == closed && closed == n && open == n {
		*combination = append(*combination, strings.Join(*stack, ""))
		return
	}
	if open < n {
		*stack = append(*stack, "(")
		backtrack(open+1, closed, n, combination, stack)
		*stack = (*stack)[:len(*stack)-1]
	}
	if open > closed {
		*stack = append(*stack, ")")
		backtrack(open, closed+1, n, combination, stack)
		*stack = (*stack)[:len(*stack)-1]
	}
}
