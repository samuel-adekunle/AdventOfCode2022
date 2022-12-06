package main

import "io/ioutil"

type LRUCache struct {
	capacity int
	cache    []rune
}

func (l *LRUCache) add(r rune) {
	if len(l.cache) == l.capacity {
		l.cache = append(l.cache[1:], r)
	} else {
		l.cache = append(l.cache, r)
	}
}

func (l *LRUCache) allDifferent() bool {
	if len(l.cache) < l.capacity {
		return false
	}

	for i := 0; i < l.capacity; i++ {
		for j := i + 1; j < l.capacity; j++ {
			if l.cache[i] == l.cache[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	fileBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	message := string(fileBytes)
	cache := LRUCache{capacity: 14, cache: []rune{}}

	for i, r := range message {
		cache.add(r)
		if cache.allDifferent() {
			println(i + 1)
			break
		}
	}
}
