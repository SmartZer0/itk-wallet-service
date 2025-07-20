package main

import (
	"fmt"
)

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

// 1
func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

// 2
func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

// 3
func (m *StringIntMap) Copy() map[string]int {
	copyMap := make(map[string]int)
	for k, v := range m.data {
		copyMap[k] = v
	}
	return copyMap
}

// 4
func (m *StringIntMap) Exists(key string) bool {
	_, exists := m.data[key]
	return exists
}

// 5
func (m *StringIntMap) Get(key string) (int, bool) {
	val, ok := m.data[key]
	return val, ok
}

func main() {
	m := NewStringIntMap()
	m.Add("one", 1)
	m.Add("two", 2)

	val, ok := m.Get("one")
	fmt.Printf("Get 'one': value = %d, exists = %v\n", val, ok)
	fmt.Println("Exists 'three':", m.Exists("three"))

	m.Remove("two")

	copy := m.Copy()
	fmt.Println("Copy of map:", copy)
}
