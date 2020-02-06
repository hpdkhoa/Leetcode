package _07_35_SearchInsertPosition

func binarySearch(nums []int, l int, r int, x int) int {
	mid := l + (r - l) / 2;
	if l > r {
		return -1
	}
	if nums[mid] == x {
		return mid
	}
	if x > nums[mid]{
		return binarySearch(nums,mid+1,r,x)
	} else {
		if x < nums[mid] {
			return binarySearch(nums,l,mid-1,x)
		}
	}
	return -1
}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if target > nums[len(nums) - 1] {
		return len(nums)
	}
	if target <= nums[0]{
		return 0
	}
	result := binarySearch(nums, 0, len(nums)-1, target)
	if result != -1 {
		return result
	} else {
		for i := 1; i < len(nums); i++ {
			if nums[i] >= target {
				return i
			}
		}
		return len(nums)
	}

}