package _06_27_RemoveElement

func removeElement(nums []int, val int) int {
	i := 0
	count := len(nums)
	for j := 0; j < len(nums); j++{
		if nums[j] != val {
			nums[i] = nums [j]
			i++
		} else {
			count--
		}
	}
	return count
}