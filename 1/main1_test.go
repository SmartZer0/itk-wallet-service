package main

import (
	"testing"
)

func TestJoinToString(t *testing.T) {
	str := joinToString(1, "Go", true, 3.14)
	expected := "1Gotrue3.14"
	if str != expected {
		t.Errorf("Ожидалась строка %s, но получено %s", expected, str)
	}
}

func TestHashWithSalt(t *testing.T) {
	runes := []rune("teststring")
	salt := "go-2024"
	hash := hashWithSalt(runes, salt)

	if len(hash) != 64 {
		t.Errorf("Длина хеша должна быть 64 символа, получено %d", len(hash))
	}
}
