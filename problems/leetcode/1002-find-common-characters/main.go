package main

import "fmt"

func main() {
	data := []string{"acabcddd", "bcbdbcbd", "baddbadb", "cbdddcac", "aacbcccd", "ccccddda", "cababaab", "addcaccd"}
	fmt.Println(commonChars(data))
}

func commonChars(words []string) []string {
	m := make(map[string]int)
	for i := 0; i < len(words[0]); i++ {
		c := string(words[0][i])
		m[c]++
	}

	for i := 1; i < len(words); i++ {
		wm := make(map[string]int)
		for j := 0; j < len(words[i]); j++ {
			c := string(words[i][j])
			wm[c]++
		}
		for k, v := range m {
			count, found := wm[k]
			if !found {
				m[k] = 0
			} else if count < v && v > 0 {
				diff := count - v
				if diff < 0 {
					diff *= -1
				}
				m[k] = v - diff
			}

		}
	}

	res := make([]string, 0)
	for k, v := range m {
		if v <= 0 {
			continue
		}
		for i := 0; i < v; i++ {
			res = append(res, k)
		}
	}
	return res
}
