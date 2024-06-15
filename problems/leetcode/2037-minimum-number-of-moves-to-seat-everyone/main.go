package main

import (
	"fmt"
	"sort"
)

func main() {
	// seats := []int{3, 1, 5}
	// students := []int{2, 7, 4}
	seats := []int{4, 1, 5, 9}
	students := []int{1, 3, 2, 6}
	// seats := []int{14, 3, 1, 1, 17, 12, 3, 19, 8}
	// students := []int{1, 11, 12, 17, 6, 1, 1, 7, 10}
	// seats := []int{1, 1, 3, 3}
	// students := []int{1, 1, 1, 1}
	fmt.Println(minMovesToSeat(seats, students))
}

// success, still slow
func minMovesToSeat(seats []int, students []int) int {
	sort.Slice(seats, func(i, j int) bool {
		return seats[i] < seats[j]
	})

	sort.Slice(students, func(i, j int) bool {
		return students[i] < students[j]
	})

	sum := 0
	sl := len(seats)
	for i := 0; i < sl; i++ {
		d := seats[i] - students[i]
		if d < 0 {
			d *= -1
		}
		sum += d
	}
	return sum
}

// success, slow
func minMovesToSeat_success_slow(seats []int, students []int) int {
	sort.Slice(seats, func(i, j int) bool {
		return seats[i] < seats[j]
	})

	sort.Slice(students, func(i, j int) bool {
		return students[i] < students[j]
	})

	m := make(map[int][]int, len(seats))
	for i := 0; i < len(seats); i++ {
		m[seats[i]] = append(m[seats[i]], i)
	}

	l := len(students)
	for i := l - 1; i >= 0; i-- {
		if arr, ok := m[students[i]]; ok {
			foundSeatIndex := arr[len(arr)-1]
			seats = append(seats[:foundSeatIndex], seats[foundSeatIndex+1:]...)

			if len(arr) > 1 {
				m[students[i]] = arr[:len(arr)-1]
			} else {
				delete(m, students[i])
			}
			students = append(students[:i], students[i+1:]...)
		}
	}

	sum := 0
	sl := len(seats)
	for i := 0; i < sl; i++ {
		d := seats[i] - students[i]
		if d < 0 {
			d *= -1
		}
		sum += d
	}
	// for i := 0; i < sl; i++ {
	// 	min := seats[i] - students[0]
	// 	if min < 0 {
	// 		min = min * -1
	// 	}
	// 	index := 0
	// 	for j := 1; j < len(students); j++ {
	// 		d := seats[i] - students[j]
	// 		if d < 0 {
	// 			d = d * -1

	// 		}

	// 		if d < min {
	// 			min = d
	// 			index = j
	// 		}
	// 	}

	// 	students = append(students[:index], students[index+1:]...)
	// 	sum += min
	// }

	return sum
}
