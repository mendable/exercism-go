package luhn

import "strconv"
import "strings"

func Valid(num string) bool {
	num = strings.TrimSpace(num)
	num = strings.Replace(num, " ", "", -1)

	if len(num) <= 1 {
		return false
	}

	var total int

	for i := len(num) - 1; i >= 0; i-- {
		c, _ := strconv.Atoi(string(num[i]))

		if i%2 == 0 {
			total += c
		} else {
			tmp := c * 2
			if tmp > 9 {
				total += (tmp - 9)
			} else {
				total += tmp
			}
		}
	}

	if total%10 == 0 {
		return true
	} else {
		return false
	}
}
