package leetcode

func canFinishKahn(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	inDegree := make([]int, numCourses)
	for _, pair := range prerequisites {
		graph[pair[1]] = append(graph[pair[1]], pair[0])
		inDegree[pair[0]]++
	}
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	count := 0
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		count++
		for _, neighbor := range graph[course] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	return count == numCourses
}
