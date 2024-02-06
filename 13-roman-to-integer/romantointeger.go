/*
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
*/
package romantointeger

func RomanToInt(s string) int {
	var total int

	for i, v := range s {

		if i < len(s)-1 {
			switch {
			case v == 'I' && (s[i+1] == 'X' || s[i+1] == 'V'):
				total = total - 1
				continue
			case v == 'X' && (s[i+1] == 'L' || s[i+1] == 'C'):
				total = total - 10
				continue
			case v == 'C' && (s[i+1] == 'D' || s[i+1] == 'M'):
				total = total - 100
				continue
			}
		}

		switch {
		case v == 'I':
			total = total + 1
			continue
		case v == 'V':
			total = total + 5
			continue
		case v == 'X':
			total = total + 10
			continue
		case v == 'L':
			total = total + 50
			continue
		case v == 'C':
			total = total + 100
			continue
		case v == 'D':
			total = total + 500
			continue
		case v == 'M':
			total = total + 1000
			continue
		}
	}

	return total

}
