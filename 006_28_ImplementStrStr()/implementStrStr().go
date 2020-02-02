package _06_28_ImplementStrStr__

func strStr(haystack string, needle string) int {
	if needle == ""{
		return 0
	}
	if len(haystack) == len(needle) {
		if haystack == needle {
			return 0
		}
		return -1
	}
	for i := 0; i < len(haystack) - len(needle) + 1; i++ {
		if len(needle) > 1 {
			if string(haystack[i:i+len(needle)]) == needle{
				return i
			}
		} else {
			if string(haystack[i]) == needle{
				return i
			}
		}
	}
	return -1
}