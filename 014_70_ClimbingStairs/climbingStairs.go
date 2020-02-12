package _14_70_ClimbingStairs

func climbStairs(n int) int {
	if n == 0{
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	first := 1
	second := 2
	for i := 3; i <= n; i++ {
		result := first + second
		first = second
		second = result
	}
	return second
}