package leetcode

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	digitToChar := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var result []string
	var backtrack func(index int, path string)
	backtrack = func(index int, path string) {
		if index == len(digits) {
			result = append(result, path)
			return
		}
		for _, char := range digitToChar[digits[index]] {
			backtrack(index+1, path+string(char))
		}
	}

	backtrack(0, "")
	return result
}
