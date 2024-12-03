func removeElement(nums []int, val int) int {
  k := 0
  for i := 0; i < len(nums); i++ { 
    if val != nums[i] {
      nums[k] = nums[i]
      k += 1
    }
  } 
  return k
}
