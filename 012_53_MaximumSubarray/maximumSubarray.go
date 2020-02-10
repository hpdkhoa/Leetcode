package _12_53_MaximumSubarray

func maxSubArray(nums []int) int {
	maxsum := nums[0]
	for i := 1; i < len(nums) ; i++ {
		if nums[i] + nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if maxsum < nums[i] {
			maxsum = nums[i]
		}
	}
	return maxsum
}