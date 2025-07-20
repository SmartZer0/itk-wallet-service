package main

import (
	"testing"
	"time"
)

func TestStartRandomGenerator(t *testing.T) {
	randCh := startRandomGenerator()

	select {
	case val := <-randCh:
		if val < 0 || val >= 1000 {
			t.Errorf("Число вне диапазона: %d", val)
		}
	case <-time.After(1 * time.Second):
		t.Errorf("Не удалось получить число из канала за 1 секунду")
	}
}
