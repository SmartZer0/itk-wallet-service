package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestMerge(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
	}()

	go func() {
		defer close(ch2)
		ch2 <- 3
		ch2 <- 4
	}()

	merged := merge(ch1, ch2)
	var result []int
	for val := range merged {
		result = append(result, val)
	}

	sort.Ints(result)
	expected := []int{1, 2, 3, 4}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}
