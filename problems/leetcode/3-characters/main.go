package main

import "fmt"

func main() {
	v := ""

	v = "abcabcbb"
	fmt.Println(3, lengthOfLongestSubstring(v))

	v = "au"
	fmt.Println(2, lengthOfLongestSubstring(v))

	v = "dvdf"
	fmt.Println(3, lengthOfLongestSubstring(v))

	v = "jbpnbwwd"
	fmt.Println(4, lengthOfLongestSubstring(v))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	m := make(map[byte]struct{})
	left, right := 0, 1
	m[s[left]] = struct{}{}
	maxLength := len(m)
	for right < len(s) {
		c := s[right]
		_, found := m[c]
		if found {
			if maxLength < len(m) {
				maxLength = len(m)
			}
			for mk := range m {
				delete(m, mk)
				left++
			}
			left--
			m[s[right]] = struct{}{}
			right = left + 1
		} else {
			m[c] = struct{}{}
			right++
		}
	}

	if maxLength < len(m) {
		maxLength = len(m)
	}

	return maxLength
}

// faster if you delete map instead of recreating it
func lengthOfLongestSubstring_answer2(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	m := make(map[byte]struct{})
	left, right := 0, 1
	m[s[left]] = struct{}{}
	maxLength := len(m)
	for right < len(s) {
		c := s[right]
		_, found := m[c]
		if found {
			if maxLength < len(m) {
				maxLength = len(m)
			}
			for mk := range m {
				delete(m, mk)
			}
			left++
			m[s[left]] = struct{}{}
			right = left + 1
		} else {
			m[c] = struct{}{}
			right++
		}
	}

	if maxLength < len(m) {
		maxLength = len(m)
	}

	return maxLength
}

func lengthOfLongestSubstring_answer1(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	m := make(map[byte]struct{})
	left, right := 0, 1
	m[s[left]] = struct{}{}
	maxLength := len(m)
	for right < len(s) {
		c := s[right]
		_, found := m[c]
		if found {
			if maxLength < len(m) {
				maxLength = len(m)
			}
			m = make(map[byte]struct{})
			left++
			m[s[left]] = struct{}{}
			right = left + 1
		} else {
			m[c] = struct{}{}
			right++
		}
	}

	if maxLength < len(m) {
		maxLength = len(m)
	}

	return maxLength
}

func lengthOfLongestSubstring_boolValueSlower(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	m := make(map[byte]bool)
	left, right := 0, 1
	m[s[left]] = true
	maxLength := len(m)
	for right < len(s) {
		c := s[right]
		_, found := m[c]
		if found {
			if maxLength < len(m) {
				maxLength = len(m)
			}
			m = make(map[byte]bool)
			left++
			m[s[left]] = true
			right = left + 1
		} else {
			m[c] = true
			right++
		}
	}

	if maxLength < len(m) {
		maxLength = len(m)
	}
	return maxLength
}

// fail: timeout
func lengthOfLongestSubstring_failTimeout(s string) int {

	l := 0
	for i := 0; i < len(s); i++ {
		m := make(map[string]struct{}, len(s[i:]))
		for _, r := range s[i:] {
			c := string(r)
			_, found := m[c]
			if found {
				if l < len(m) {
					l = len(m)
				}
				m = make(map[string]struct{})
			}
			m[c] = struct{}{}
		}
		if l < len(m) {
			l = len(m)
		}
	}

	return l
}
