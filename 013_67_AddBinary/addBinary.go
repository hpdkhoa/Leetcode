package _13_67_AddBinary

import "strconv"

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len (b) - 1
	count := 0
	res := ""
	for {
		if i < 0 && j < 0 {
			if count == 1 {
				res = strconv.FormatInt(int64(count),10) + res
			}
			return res
		}
		if i >= 0 && j >= 0 {
			count += int(a[i] - '0') + int(b[j] - '0')
			if count == 2 {
				count = 1
				res = "0" + res
			} else {
				if count == 3 {
					count = 1
					res = "1" + res
				} else {
					res = strconv.FormatInt(int64(count),10) + res
					count = 0
				}
			}
			i--
			j--
		} else {
			if i < 0 {
				count += int(b[j] - '0')
				if count == 2 {
					count = 1
					res = "0" + res
				} else {
					res = strconv.FormatInt(int64(count),10) + res
					count = 0
				}
				j--
			} else {
				count += int(a[i] - '0')
				if count == 2 {
					count = 1
					res = "0" + res
				} else {
					res = strconv.FormatInt(int64(count),10) + res
					count = 0
				}
				i--
			}
		}
	}
}