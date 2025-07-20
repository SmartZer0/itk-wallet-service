package main

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	s1 := []string{"a", "b", "c", "d"}
	s2 := []string{"b", "d", "e"}
	expected := []string{"a", "c"}

	result := difference(s1, s2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

func TestDifferenceEmptySecond(t *testing.T) {
	s1 := []string{"x", "y"}
	s2 := []string{}
	expected := []string{"x", "y"}

	result := difference(s1, s2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

func TestDifferenceEmptyFirst(t *testing.T) {
	s1 := []string{}
	s2 := []string{"x", "y"}
	expected := []string{}

	result := difference(s1, s2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

func TestDifferenceNoDifference(t *testing.T) {
	s1 := []string{"m", "n"}
	s2 := []string{"m", "n"}
	expected := []string{}

	result := difference(s1, s2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}
