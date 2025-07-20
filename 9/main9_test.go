package main

import (
	"reflect"
	"testing"
)

func TestCubePipeline(t *testing.T) {
	in := make(chan uint8)
	out := cubePipeline(in)

	go func() {
		defer close(in)
		for _, v := range []uint8{2, 4, 6} {
			in <- v
		}
	}()

	var result []float64
	for v := range out {
		result = append(result, v)
	}

	expected := []float64{8, 64, 216}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}
