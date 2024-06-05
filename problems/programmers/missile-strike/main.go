package main

import (
	"fmt"
)

func main() {
	data := [][]int{{4, 5}, {4, 8}, {10, 14}, {11, 13}, {5, 12}, {3, 7}, {1, 4}}
	// data := [][]int{{4, 5}, {3, 7}, {1, 4}}
	// data := [][]int{{4, 5}, {4, 8}, {3, 7}}
	res := f2(data)
	if res == 3 {
		fmt.Println("ok", res)
		return
	}
	fmt.Println("wrong", res)
}

func f2(targets [][]int) int {
	targets = mergeSort(targets)
	fmt.Println(targets)

	// m := make(map[int][]int)
	fired := []int{targets[0][1] - 1}
	for i := 1; i < len(targets); i++ {
		start := targets[i][0]
		end := targets[i][1]
		lastFired := fired[len(fired)-1]
		if lastFired < start || lastFired >= end {
			fired = append(fired, targets[i][1]-1)
		}
		// if lastFired >= start && lastFired < end {
		// 	// fired = append(fired, targets[i][1])
		// 	fmt.Println("skip")
		// } else {
		// 	fired = append(fired, targets[i][1])
		// }
	}
	fmt.Println(targets)
	fmt.Println(fired)
	return 0
}

func sortTargets(targets [][]int) [][]int {
	for i := 0; i < len(targets)-1; i++ {
		for j := i + 1; j < len(targets); j++ {
			if targets[i][1] > targets[j][1] {
				tmp := targets[i]
				targets[i], targets[j] = targets[j], tmp
			}
		}
	}
	return targets
}

func merge(left, right [][]int) [][]int {
	result := make([][]int, 0)
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		v := make([]int, 2)
		if left[i][1] <= right[j][1] {
			copy(v, left[i])
			result = append(result, v)
			i++
		} else {
			copy(v, right[j])
			result = append(result, v)
			j++
		}
	}

	for k := i; k < len(left); k++ {
		v := make([]int, 2)
		copy(v, left[k])
		result = append(result, v)
	}
	for k := j; k < len(right); k++ {
		v := make([]int, 2)
		copy(v, right[k])
		result = append(result, v)
	}
	return result
}

func mergeSort(target [][]int) [][]int {
	if len(target) <= 1 {
		return target
	}

	mid := len(target) / 2
	left := mergeSort(target[:mid])
	right := mergeSort(target[mid:])
	return merge(left, right)
}

func f(data [][]int) int {
	numMissiles := len(data)
	points := make(map[int][]int)
	overlap := make(map[int][]int)
	for i := 0; i < numMissiles; i++ {
		points[data[i][0]] = []int{}
		// firePoints[data[i][0]] = make(map[int]struct{})

		overlap[i] = []int{}
	}

	for i := 0; i < numMissiles; i++ {
		is := data[i][0]
		ie := data[i][1]
		points[is] = add(points[is], i)
		for j := i + 1; j < numMissiles; j++ {
			js := data[j][0]
			je := data[j][1]

			if is <= js && ie > js {
				overlap[i] = append(overlap[i], j)
				overlap[j] = append(overlap[j], i)
				points[js] = add(points[js], i)
				points[js] = add(points[js], j)
			} else if js <= is && is < je {
				overlap[i] = append(overlap[i], j)
				overlap[j] = append(overlap[j], i)
				points[is] = add(points[is], i)
				points[is] = add(points[is], j)
			}
		}
	}

	// for k, v := range overlap {

	// }

	fmt.Println("overlap", overlap)
	fmt.Println("firePoints", points)
	return len(points)
}

func add(points []int, point int) []int {
	for i := range points {
		if point == points[i] {
			return points
		}
	}
	res := append(points, point)
	return res
}
