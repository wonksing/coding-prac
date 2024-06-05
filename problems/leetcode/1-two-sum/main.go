package main

import "fmt"

func main() {
	var nums []int
	// nums := []int{2, 7, 11, 15}
	// fmt.Println(twoSum(nums, 9))

	nums = []int{2, 7, 11, 15}
	fmt.Println(twoSumPerf(nums, 9))

	nums = []int{3, 2, 4}
	fmt.Println(twoSumPerf(nums, 6))

	nums = []int{2, 7, 11, 15}
	fmt.Println(twoSumPerf2(nums, 9))

	nums = []int{3, 2, 4}
	fmt.Println(twoSumPerf2(nums, 6))
}

// twoSum
func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums)-1; i++ {
		a := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if a+nums[j] == target {
				fmt.Println("found", a, nums[j])
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSumPerf(nums []int, target int) []int {
	m := make(map[int]int)
	m[0] = nums[0]

	i := 0
	j := i + 1
	for {
		a := m[i]
		b := nums[j]
		if a+b == target {
			return []int{i, j}
		}

		if i+1 == len(nums)-1 {
			break
		}

		if len(nums)-1 <= j {
			i++
			j = i + 1
			m[i] = nums[i]
		} else {
			j++
		}
	}
	return []int{}
}

func twoSumPerf2(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		a := nums[i]
		v := target - a
		if _, ok := m[v]; ok {
			return []int{i, m[v]}
		}
		m[a] = i
	}
	return []int{}
}
