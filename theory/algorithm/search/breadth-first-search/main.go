package main

import "fmt"

func main() {
	graph := map[int][]int{
		1: {2, 3},
		2: {4, 5},
		3: {},
		4: {},
		5: {},
	}

	fmt.Println(bfs(graph, 1, 3))

}

func bfs(graph map[int][]int, start, target int) bool {
	visited := make(map[int]bool)
	queue := make([]int, 0)

	queue = append(queue, start)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == target {
			return true
		}

		if !visited[node] {
			visited[node] = true

			for _, neighbor := range graph[node] {
				if !visited[neighbor] {
					queue = append(queue, neighbor)
				}
			}
		}
	}

	return false
}
