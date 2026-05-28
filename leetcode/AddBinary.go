package leetcode

func addBinary(a string, b string) string {
	var result []byte
	i := len(a) - 1
	j := len(b) - 1
	carry := 0

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}

		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		result = append(result, byte(sum%2+'0'))
		carry = sum / 2
	}

	for k := 0; k < len(result)/2; k++ {
		result[k], result[len(result)-1-k] = result[len(result)-1-k], result[k]
	}

	return string(result)
}
