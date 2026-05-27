package leetcode

func isMatch(stringToMatch string, pattern string) bool {
	stringPointer, patternPointer := 0, 0
	stringPositionMatchedByStar := 0
	lastStarPatternIndex := -1

	for stringPointer < len(stringToMatch) {
		if patternPointer < len(pattern) && (pattern[patternPointer] == stringToMatch[stringPointer] || pattern[patternPointer] == '?') {
			stringPointer++
			patternPointer++
		} else if patternPointer < len(pattern) && pattern[patternPointer] == '*' {
			lastStarPatternIndex = patternPointer
			stringPositionMatchedByStar = stringPointer
			patternPointer++
		} else if lastStarPatternIndex != -1 {
			patternPointer = lastStarPatternIndex + 1
			stringPositionMatchedByStar++
			stringPointer = stringPositionMatchedByStar
		} else {
			return false
		}
	}

	for patternPointer < len(pattern) && pattern[patternPointer] == '*' {
		patternPointer++
	}

	return patternPointer == len(pattern)
}
