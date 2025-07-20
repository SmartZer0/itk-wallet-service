package main

import (
	"fmt"
)

// - true, если есть хотя бы одно общее значение между a и b
// - слайс общих значений без повторов
func findIntersection(a, b []int) (bool, []int) {
	setA := make(map[int]struct{})
	for _, val := range a {
		setA[val] = struct{}{}
	}

	seen := make(map[int]struct{})
	var intersection []int
	for _, val := range b {
		if _, found := setA[val]; found {
			if _, alreadyAdded := seen[val]; !alreadyAdded {
				intersection = append(intersection, val)
				seen[val] = struct{}{}
			}
		}
	}

	return len(intersection) > 0, intersection
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	hasIntersection, common := findIntersection(a, b)
	fmt.Println("Есть пересечения?", hasIntersection)
	fmt.Println("Общие значения:", common)
}
