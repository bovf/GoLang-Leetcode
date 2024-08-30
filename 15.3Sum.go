package golangleetcode

import (
	"fmt"
	"sort"
)

// O(n^3) - Too slow
// func removeDuplicateInt(intSlice []int) []int {
// 	allKeys := make(map[int]int)
// 	list := []int{}
//
// 	for _, item := range intSlice {
// 		if value := allKeys[item]; value < 4 {
// 			allKeys[item] += 1
// 			list = append(list, item)
// 		}
// 	}
// 	return list
// }
//
// func threeSum(nums []int) [][]int {
// 	result := [][]int{}
// 	resultMap := make(map[string]bool)
// 	nums = removeDuplicateInt(nums)
//
// 	for i := 0; i < len(nums); i++ {
// 		for n := i + 1; n < len(nums); n++ {
// 			for z := n + 1; z < len(nums); z++ {
// 				if nums[i]+nums[n]+nums[z] == 0 {
// 					triplets := []int{nums[i], nums[n], nums[z]}
// 					sort.Ints(triplets)
// 					mapString := fmt.Sprintf("%d,%d,%d", triplets[0], triplets[1], triplets[2])
// 					if !resultMap[mapString] {
// 						result = append(result, []int{nums[i], nums[n], nums[z]})
// 						resultMap[mapString] = true
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return result
// }

func threeSum(nums []int) [][]int {
	result := [][]int{}
	
	sort.Ints(nums)

	for i := 0; i < len(nums) - 2; i++ {
		
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i + 1, len(nums) - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left + 1] {
					left++
				}

				for left < right && nums[right] == nums[right - 1] {
					right--
				}

				left++
				right--
			} else if sum < 0 {
					left++
			} else {
					right--
			}
		}
	}
	return result
}
