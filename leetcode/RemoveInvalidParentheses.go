package leetcode

func removeInvalidParentheses(s string) []string {
	var result []string
	var queue []string
	visited := make(map[string]bool)
	queue = append(queue, s)
	visited[s] = true
	found := false
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if isValid(current) {
			result = append(result, current)
			found = true
		}
		if found {
			continue
		}
		for i := 0; i < len(current); i++ {
			if current[i] != '(' && current[i] != ')' {
				continue
			}
			next := current[:i] + current[i+1:]
			if !visited[next] {
				queue = append(queue, next)
				visited[next] = true
			}
		}
	}
	return result
}

func isValid(s string) bool {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			count++
		} else if s[i] == ')' {
			count--
			if count < 0 {
				return false
			}
		}
	}
	return count == 0
}
