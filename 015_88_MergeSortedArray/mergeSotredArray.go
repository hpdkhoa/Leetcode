package _15_88_MergeSortedArray

func swap(x, y int) (int,int) {
	return y,x
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}
	for i := 0; i < n+m; i++ {
		for j := i; j < n+m; j++ {
			if nums1[i] > nums1[j] {
				for x := j; x > i; x-- {
					nums1[x], nums1[x-1] = swap (nums1[x], nums1[x-1])
				}
			}
		}
	}
}