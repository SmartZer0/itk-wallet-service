package main

import (
	"fmt"
)

func difference(slice1, slice2 []string) []string {
	set2 := make(map[string]struct{})
	for _, s := range slice2 {
		set2[s] = struct{}{}
	}

	var result []string
	for _, s := range slice1 {
		if _, found := set2[s]; !found {
			result = append(result, s)
		}
	}

	if result == nil {
		return []string{}
	}
	return result

}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	diff := difference(slice1, slice2)
	fmt.Println("Результат:", diff)
}
