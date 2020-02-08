package _10_58_LengthOfLastWord

import "strings"

func lengthOfLastWord(s string) int {
	strs := strings.Split(s, " ")
	if len(strs) == 0 {
		return 0
	}
	for i := len(strs) - 1; i >= 0; i--{
		if len(strs[i]) > 0{
			return len(strs[i])
		}
	}
	return 0
}