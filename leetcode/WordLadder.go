package leetcode

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, w := range wordList {
		wordSet[w] = true
	}

	if !wordSet[endWord] {
		return 0
	}

	queue := []string{beginWord}
	steps := 1

	for len(queue) > 0 {
		for size := len(queue); size > 0; size-- {
			word := queue[0]
			queue = queue[1:]

			if word == endWord {
				return steps
			}

			b := []byte(word)
			for i := range b {
				orig := b[i]
				for c := byte('a'); c <= 'z'; c++ {
					if c == orig {
						continue
					}
					b[i] = c
					next := string(b)
					if wordSet[next] {
						queue = append(queue, next)
						delete(wordSet, next)
					}
				}
				b[i] = orig
			}
		}
		steps++
	}

	return 0
}
