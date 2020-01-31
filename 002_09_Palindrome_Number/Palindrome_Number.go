package _02_09_Palindrome_Number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tail := 0
	for head := x; head != 0;{
		tail = tail * 10 + head % 10
		head /= 10
	}
	return tail == x
}