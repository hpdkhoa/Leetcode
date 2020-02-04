package _05_26_RemoveDuplicatefromSortedArray

func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] == nums[i+1]{
			for j := i; j < len(nums) - 1; j++{
				nums[j] = nums[j+1]
			}
			nums[len(nums)-1] = 0
			nums = nums[:len(nums)-1]
			i--
		}
	}
	return len(nums)
}