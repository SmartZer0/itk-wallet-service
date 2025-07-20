package main

import (
	"reflect"
	"testing"
)

func TestAddAndGet(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 10)

	val, ok := m.Get("a")
	if !ok || val != 10 {
		t.Errorf("Ожидалось 10, получено %d", val)
	}
}

func TestRemove(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 10)
	m.Remove("a")

	_, ok := m.Get("a")
	if ok {
		t.Errorf("Ключ 'a' должен быть удалён")
	}
}

func TestExists(t *testing.T) {
	m := NewStringIntMap()
	m.Add("b", 20)

	if !m.Exists("b") {
		t.Errorf("Ключ 'b' должен существовать")
	}
	if m.Exists("z") {
		t.Errorf("Ключ 'z' не должен существовать")
	}
}

func TestCopy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("x", 5)
	m.Add("y", 6)

	copied := m.Copy()

	if !reflect.DeepEqual(copied, m.data) {
		t.Errorf("Копия карты не совпадает с оригиналом")
	}

	m.Add("x", 100)
	if copied["x"] == 100 {
		t.Errorf("Копия карты изменилась при изменении оригинала")
	}
}
