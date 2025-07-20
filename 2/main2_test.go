package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4}
	result := sliceExample(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

func TestAddElements(t *testing.T) {
	input := []int{1, 2, 3}
	result := addElements(input, 4)
	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

func TestCopySlice(t *testing.T) {
	input := []int{1, 2, 3}
	copy := copySlice(input)
	input[0] = 99
	if reflect.DeepEqual(input, copy) {
		t.Errorf("Копия не должна меняться при изменении оригинала")
	}
}

func TestRemoveElement(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	result := removeElement(input, 2)
	expected := []int{1, 2, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

func TestRemoveElementOutOfBounds(t *testing.T) {
	input := []int{1, 2, 3}
	result := removeElement(input, 5)
	if !reflect.DeepEqual(result, input) {
		t.Errorf("Если индекс вне границ, слайс не должен измениться")
	}
}
