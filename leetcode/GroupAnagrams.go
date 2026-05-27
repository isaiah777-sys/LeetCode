package leetcode

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)
	for _, str := range strs {
		sortedStr := sortString(str)
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}
	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}
	return result
}

func sortString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)-1; i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}
