package leetcode

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}

	var result float64 = 1.0
	currentProduct := x

	for n > 0 {
		if n%2 == 1 {
			result *= currentProduct
		}
		currentProduct *= currentProduct
		n /= 2
	}

	return result
}
