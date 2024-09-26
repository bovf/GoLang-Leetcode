package main

import (
	"math"
	"sort"
)

// func threeSumClosest(nums []int, target int) int {
// 	sort.Ints(nums)
// 	result := nums[0] + nums[1] + nums[2]
//
// 	for i := 0; i < len(nums)-2; i++ {
//
// 		left, right := i+1, len(nums)-1
// 		for left < right {
// 			sum := nums[i] + nums[left] + nums[right]
// 			distance := sum - target
// 			old_distance := result - target
//
// 			if math.Abs(float64(distance)) < math.Abs(float64(old_distance)) {
// 				result = sum
// 			}
//
// 			if sum < target {
// 				left++
// 			} else if sum > target {
// 				right--
// 			} else {
// 				return sum
// 			}
// 		}
// 	}
//
// 	return result
// }

// Faster
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closestSum := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {

		left, right := i+1, len(nums)-1
		for left < right {
			currentSum := nums[i] + nums[left] + nums[right]

			if currentSum == target {
				return target
			} else if currentSum < target {
				left++
			} else {
				right--
			}
			if math.Abs(float64(currentSum-target)) < math.Abs(float64(closestSum-target)) {
				closestSum = currentSum
			}
		}
	}

	return closestSum
}
