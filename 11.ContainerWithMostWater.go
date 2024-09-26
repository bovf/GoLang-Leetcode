package main

// Very slow and bruteforce
// func maxArea(height []int) int {
// 	result := 0
// 	for i := 0; i < len(height); i++ {
// 		for x := i + 1; x < len(height); x++ {
// 			area := min(height[i], height[x]) * (x - i)
// 			if area > result {
// 				result = area
// 			}
// 		}
// 	}
// 	return result
// }

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	result := 0
	for left < right {
		width := right - left
		area := min(height[left], height[right]) * width

		if area > result {
			result = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return result
}
