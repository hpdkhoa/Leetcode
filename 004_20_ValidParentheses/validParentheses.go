package _04_20_ValidParentheses

func isValid(s string) bool {
	substr := ""
	m := make(map[string]string)
	m[")"] = "("
	m["]"] = "["
	m["}"] = "{"

	for _, chr := range s{
		if m[string(chr)] != "" {
			if substr == "" {
				return false
			}
			if string(substr[len(substr)-1]) != m[string(chr)]{
				return false
			} else {
				substr = substr[:len(substr)-1]
			}
		} else {
			substr = substr + string(chr)
		}
	}
	if substr != "" {
		return false
	}
	return true
}