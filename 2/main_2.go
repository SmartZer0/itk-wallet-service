package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	// 1
	originalSlice := generateRandomSlice(10)
	fmt.Println("Original slice:", originalSlice)

	// 2
	evenSlice := sliceExample(originalSlice)
	fmt.Println("Even slice:", evenSlice)

	// 3
	extendedSlice := addElements(originalSlice, 42)
	fmt.Println("Slice after adding 42:", extendedSlice)

	// 4
	copied := copySlice(originalSlice)
	originalSlice[0] = -1
	fmt.Println("Original after modification:", originalSlice)
	fmt.Println("Copied slice:", copied)

	// 5
	indexToRemove := 3
	removed := removeElement(originalSlice, indexToRemove)
	fmt.Printf("Slice after removing index %d: %v\n", indexToRemove, removed)
}

func generateRandomSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = rand.Intn(100)
	}
	return s
}

func sliceExample(slice []int) []int {
	var result []int
	for _, v := range slice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

func addElements(slice []int, elem int) []int {
	return append(slice, elem)
}

func copySlice(slice []int) []int {
	copied := make([]int, len(slice))
	copy(copied, slice)
	return copied
}

func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
