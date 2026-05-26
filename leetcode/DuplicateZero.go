package leetcode

func duplicateZeros(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			for j := n - 1; j > i; j-- {
				arr[j] = arr[j-1]
			}
			if i+1 < n {
				arr[i+1] = 0
			}
			i++
		}
	}
	return arr
}
