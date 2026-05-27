package leetcode

func findSubstring(s string, words []string) []int {
	if len(words) == 0 || len(s) == 0 {
		return []int{}
	}
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}
	wordLength := len(words[0])
	substringLength := wordLength * len(words)
	result := []int{}
	for i := 0; i <= len(s)-substringLength; i++ {
		seen := make(map[string]int)
		for j := 0; j < len(words); j++ {
			word := s[i+j*wordLength : i+(j+1)*wordLength]
			if count, exists := wordCount[word]; !exists || seen[word] >= count {
				break
			}
			seen[word]++
			if j == len(words)-1 {
				result = append(result, i)
			}
		}
	}
	return result
}

//Alternative approach to solve the problem

func findSubstringAlternative(s string, words []string) []int {
	if len(words) == 0 || len(s) == 0 {
		return []int{}
	}
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}
	wordLength := len(words[0])
	substringLength := wordLength * len(words)
	result := []int{}
	for i := 0; i < wordLength; i++ {
		for j := i; j <= len(s)-substringLength; j += wordLength {
			seen := make(map[string]int)
			for k := 0; k < len(words); k++ {
				word := s[j+k*wordLength : j+(k+1)*wordLength]
				if count, exists := wordCount[word]; !exists || seen[word] >= count {
					break
				}
				seen[word]++
				if k == len(words)-1 {
					result = append(result, j)
				}
			}
		}
	}
	return result
}
