package golangleetcode

import (
	"math"
)

func longestCommonPrefix(strs []string) string {
	result := ""
	shortestString := math.MaxInt64
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < shortestString {
			shortestString = len(strs[i])
		}
	}
	for i := 0; i < shortestString; i++ {
		for n := 0; n < len(strs); n++ {
			if strs[n][i] != strs[0][i] {
				return result
			}
		}
		result += string(strs[0][i])
	}
	return result
}
