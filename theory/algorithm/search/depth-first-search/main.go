package main

import "fmt"

func main() {
	// graph := make(map[int][]int)
	graph := map[int][]int{
		1: {2, 3},
		2: {4, 5},
		3: {},
		4: {},
		5: {},
	}
	visited := make(map[int]bool)
	fmt.Println(dfs(graph, 1, 3, visited))
	fmt.Println(visited)
}

func dfs(graph map[int][]int, start, target int, visited map[int]bool) bool {
	if start == target {
		return true
	}

	visited[start] = true
	for _, neighbor := range graph[start] {
		if !visited[neighbor] {
			if dfs(graph, neighbor, target, visited) {
				return true
			}
		}
	}

	return false
}
