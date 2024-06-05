package main

import "fmt"

func main() {
	hashMap := NewHashMap(10)
	hashMap.Put("name", "Alice")
	hashMap.Put("age", "30")
	hashMap.Put("city", "New York")

	// 값 가져오기
	if value, found := hashMap.Get("name"); found {
		fmt.Println("name:", value)
	} else {
		fmt.Println("name not found")
	}

	// 값 제거
	hashMap.Remove("age")

	// 제거 후 값 가져오기
	if value, found := hashMap.Get("age"); found {
		fmt.Println("age:", value)
	} else {
		fmt.Println("age not found")
	}
}

type Entry struct {
	key   string
	value string
}

type HashMap struct {
	buckets [][]Entry
	size    int
}

func NewHashMap(size int) *HashMap {
	return &HashMap{
		buckets: make([][]Entry, size),
		size:    size,
	}
}

func (h *HashMap) hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		// 31은 소수이고 충돌을 최소화 하기 위함
		hash = (31 * hash) + int(key[i])
	}

	// buckets 사이즈 안에서 처리하기 위함
	return hash % h.size
}

func (h *HashMap) Put(key, value string) {
	index := h.hash(key)
	for i, entry := range h.buckets[index] {
		if entry.key == key {
			h.buckets[index][i].value = value
			return
		}
	}
	h.buckets[index] = append(h.buckets[index], Entry{key, value})
}

func (h *HashMap) Get(key string) (string, bool) {
	index := h.hash(key)
	for _, entry := range h.buckets[index] {
		if entry.key == key {
			return entry.value, true
		}
	}
	return "", false
}

func (h *HashMap) Remove(key string) {
	index := h.hash(key)
	for i, entry := range h.buckets[index] {
		if entry.key == key {
			h.buckets[index] = append(h.buckets[index][:i], h.buckets[index][i+1:]...)
			return
		}
	}
}
