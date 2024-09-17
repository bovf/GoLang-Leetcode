package golangleetcode

import (
	"sort"
)

func ksum(k int, start int, target int, nums *[]int, result *[][]int, quads *[]int) {
	if k != 2 {
		for i := start; i < len(*nums)-k+1; i++ {
			if i > start && (*nums)[i] == (*nums)[i-1] {
				continue
			}
			*quads = append(*quads, (*nums)[i])
			ksum(k-1, i+1, target-(*nums)[i], nums, result, quads)
			if len(*quads) > 0 {
				*quads = (*quads)[:len(*quads)-1]
			}
		}
		return
	}
	left, right := start, len(*nums)-1
	for left < right {
		if (*nums)[left]+(*nums)[right] < target {
			left++
		} else if (*nums)[left]+(*nums)[right] > target {
			right--
		} else {
			*result = append(*result, append([]int{}, append(*quads, (*nums)[left], (*nums)[right])...))
			left++
			right--
			for left < right && (*nums)[left] == (*nums)[left-1] {
				left++
			}
			for left < right && (*nums)[right] == (*nums)[right+1] {
				right--
			}
		}
	}
}

func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	quads := []int{}
	sort.Ints(nums)

	ksum(4, 0, target, &nums, &result, &quads)

	return result
}
