package main

import "fmt"

func main() {
	input := [][]int{
		0: {1, 4},
		1: {3, 4},
		2: {3, 10},
	}
	fmt.Println(solution(input))
}

func solution(v [][]int) []int {
	answer := []int{}

	x := make(map[int]int)
	y := make(map[int]int)
	for _, value := range v {
		x[value[0]]++
		y[value[1]]++
	}

	foundX, _ := findPosition(x)
	foundY, _ := findPosition(y)

	answer = append(answer, foundX, foundY)

	return answer
}

func findPosition(input map[int]int) (int, bool) {
	for key, value := range input {
		if value == 1 {
			return key, true
		}
	}
	return 0, false
}
