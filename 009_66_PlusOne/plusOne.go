package _09_66_PlusOne

func plusOne(digits []int) []int {
	for i := len(digits) - 1 ; ; {
		if i == len(digits) - 1 {
			if digits[i] != 9 {
				digits[i]++
				return digits
			} else {
				if i == 0 {
					digits[i] = 1
					digits = append(digits, 0)
					return digits
				} else {
					digits[i] = 0
					i--;
					digits[i]++
					if digits[i] != 10{
						return digits
					}
				}
			}
		}
		if digits[i] == 10 {
			if i == 0 {
				digits[i] = 1
				digits = append(digits, 0)
				return digits
			} else {
				digits[i] = 0
				i--;
				digits[i]++
			}
		} else {
			return digits
		}
	}
}