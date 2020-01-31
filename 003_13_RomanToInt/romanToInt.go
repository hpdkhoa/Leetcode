package _03_13_RomanToInt

import "strings"

func singleRomanToInt(c string) int {
	switch strings.ToUpper(c){
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		return 0
	}
}

func romanToInt(s string) int {
	result := 0
	for i := 0; i < len(s); i++{
		currentInt := singleRomanToInt(string(s[i]))
		if i != len(s) - 1 {
			nextInt := singleRomanToInt(string(s[i+1]))
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