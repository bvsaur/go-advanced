package main

import (
	"fmt"
	"time"
)

type Memory struct {
	f     Function
	cache map[int]int
}

type Function func(n int, m *Memory) int

func NewMemory(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]int),
	}
}

func (m *Memory) Get(key int) int {
	value, exists := m.cache[key]
	if !exists {
		value = m.f(key, m)
		m.cache[key] = value
	}
	return value
}

func Fibonacci(n int, m *Memory) int {
	if n <= 1 {
		return n
	}

	return m.Get(n-1) + m.Get(n-2)
}

func main() {
	cache := NewMemory(Fibonacci)
	values := []int{30, 32, 50, 40, 5, 8, 10}

	for _, key := range values {
		start := time.Now()
		result := cache.Get(key)
		fmt.Printf("Key %d, Time: %s, Result: %d\n", key, time.Since(start), result)
	}
}
