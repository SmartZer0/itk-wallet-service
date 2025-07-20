package main

import (
	"reflect"
	"testing"
)

func TestFindIntersection(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{3, 4, 5}
	expected := []int{3, 4}

	ok, result := findIntersection(a, b)
	if !ok {
		t.Errorf("Ожидалось пересечение, но его нет")
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

func TestNoIntersection(t *testing.T) {
	a := []int{10, 20}
	b := []int{30, 40}

	ok, result := findIntersection(a, b)
	if ok {
		t.Errorf("Пересечений быть не должно")
	}
	if len(result) != 0 {
		t.Errorf("Ожидался пустой слайс, получено %v", result)
	}
}

func TestWithDuplicates(t *testing.T) {
	a := []int{1, 2, 2, 3, 4}
	b := []int{2, 2, 4, 4}
	expected := []int{2, 4}

	ok, result := findIntersection(a, b)
	if !ok {
		t.Errorf("Ожидалось пересечение")
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

func TestEmptyInput(t *testing.T) {
	a := []int{}
	b := []int{}

	ok, result := findIntersection(a, b)
	if ok || len(result) != 0 {
		t.Errorf("Для пустых слайсов должно быть false и []")
	}
}
