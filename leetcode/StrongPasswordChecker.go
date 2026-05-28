package leetcode

func strongPasswordChecker(password string) int {
	hasLower, hasUpper, hasDigit := 0, 0, 0
	for i := 0; i < len(password); i++ {
		if password[i] >= 'a' && password[i] <= 'z' {
			hasLower = 1
		} else if password[i] >= 'A' && password[i] <= 'Z' {
			hasUpper = 1
		} else if password[i] >= '0' && password[i] <= '9' {
			hasDigit = 1
		}
	}
	missingTypes := 3 - (hasLower + hasUpper + hasDigit)

	if len(password) < 6 {
		return max(6-len(password), missingTypes)
	}

	var repeatingGroups []int
	for i := 0; i < len(password); {
		j := i
		for j < len(password) && password[j] == password[i] {
			j++
		}
		length := j - i
		if length >= 3 {
			repeatingGroups = append(repeatingGroups, length)
		}
		i = j
	}

	if len(password) <= 20 {
		substitutionsNeeded := 0
		for _, length := range repeatingGroups {
			substitutionsNeeded += length / 3
		}
		return max(substitutionsNeeded, missingTypes)
	}

	deletionsNeeded := len(password) - 20

	for i := range repeatingGroups {
		if deletionsNeeded > 0 && repeatingGroups[i]%3 == 0 {
			repeatingGroups[i] -= 1
			deletionsNeeded -= 1
		}
	}

	for i := range repeatingGroups {
		if deletionsNeeded >= 2 && repeatingGroups[i]%3 == 1 {
			repeatingGroups[i] -= 2
			deletionsNeeded -= 2
		}
	}

	for i := range repeatingGroups {
		if deletionsNeeded > 0 && repeatingGroups[i] >= 3 {
			removableFromGroup := repeatingGroups[i] - 2
			usedDeletions := min(deletionsNeeded, removableFromGroup-(removableFromGroup%3))

			if usedDeletions == 0 && deletionsNeeded > 0 {
				usedDeletions = min(deletionsNeeded, repeatingGroups[i]-2)
			}

			repeatingGroups[i] -= usedDeletions
			deletionsNeeded -= usedDeletions
		}
	}

	substitutionsNeeded := 0
	for _, length := range repeatingGroups {
		if length >= 3 {
			substitutionsNeeded += length / 3
		}
	}

	return (len(password) - 20) + max(substitutionsNeeded, missingTypes)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
