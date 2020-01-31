package _03_13_RomanToInt

func romanToInt(s string) int {
	m := make(map[string]int)
	m["I"] = 1
	m["V"] = 5
	m["X"] = 10
	m["L"] = 50
	m["C"] = 100
	m["D"] = 500
	m["M"] = 1000
	result := 0
	for i := 0; i < len(s); i++{
		currentInt := m[string(s[i])]
		if i != len(s) - 1 {
			nextInt := m[string(s[i+1])]
			if currentInt >= nextInt{
				result += currentInt
			} else {
				result -= currentInt - nextInt
				i++
			}
		} else {
			result += currentInt
		}
	}
	return result
}