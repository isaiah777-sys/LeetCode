package leetcode

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordSet := make(map[string]bool)
	for _, w := range wordList {
		wordSet[w] = true
	}
	if !wordSet[endWord] {
		return [][]string{}
	}

	parents := make(map[string][]string)
	distance := make(map[string]int)
	distance[beginWord] = 1

	queue := []string{beginWord}
	found := false

	for len(queue) > 0 && !found {
		levelSize := len(queue)
		localVisited := make(map[string]bool)

		for i := 0; i < levelSize; i++ {
			curr := queue[0]
			queue = queue[1:]

			currDist := distance[curr]
			wordBytes := []byte(curr)

			for j := 0; j < len(wordBytes); j++ {
				oldChar := wordBytes[j]
				for c := 'a'; c <= 'z'; c++ {
					if byte(c) == oldChar {
						continue
					}
					wordBytes[j] = byte(c)
					nextWord := string(wordBytes)

					if wordSet[nextWord] {
						if nextDist, exists := distance[nextWord]; !exists || currDist+1 <= nextDist {
							if !exists {
								distance[nextWord] = currDist + 1
								localVisited[nextWord] = true
							}
							parents[nextWord] = append(parents[nextWord], curr)
						}
						if nextWord == endWord {
							found = true
						}
					}
				}
				wordBytes[j] = oldChar
			}
		}

		for w := range localVisited {
			queue = append(queue, w)
		}
	}

	var results [][]string
	if !found {
		return results
	}

	var path []string
	var dfs func(word string)

	dfs = func(word string) {
		if word == beginWord {
			correctPath := make([]string, len(path)+1)
			correctPath[0] = beginWord
			for i, v := range path {
				correctPath[len(path)-i] = v
			}
			results = append(results, correctPath)
			return
		}

		path = append(path, word)
		for _, p := range parents[word] {
			dfs(p)
		}
		path = path[:len(path)-1]
	}

	dfs(endWord)
	return results
}
