package leetcode

func dailyTemperatures(temperatures []int) []int {
	results := make([]int, len(temperatures))
	var indexStack []int

	for currentIndex := 0; currentIndex < len(temperatures); currentIndex++ {
		for len(indexStack) > 0 && temperatures[currentIndex] > temperatures[indexStack[len(indexStack)-1]] {
			previousIndex := indexStack[len(indexStack)-1]
			indexStack = indexStack[:len(indexStack)-1]

			results[previousIndex] = currentIndex - previousIndex
		}
		indexStack = append(indexStack, currentIndex)
	}

	return results
}
