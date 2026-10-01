package cache

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestJankenSessionLifecycle(t *testing.T) {
	if !StartJanken(1, 10, "ぐー") {
		t.Fatal("start failed")
	}
	if StartJanken(1, 10, "ぱー") {
		t.Fatal("duplicate start accepted")
	}
	if _, ok := PlayJanken(1, 20, true); ok {
		t.Fatal("wrong channel accepted")
	}
	if answer, ok := PlayJanken(1, 10, false); !ok || answer != "ぐー" {
		t.Fatal("invalid move consumed session")
	}
	if answer, ok := PlayJanken(1, 10, true); !ok || answer != "ぐー" {
		t.Fatal("valid move failed")
	}
	if _, ok := PlayJanken(1, 10, true); ok {
		t.Fatal("session consumed twice")
	}
}

func TestConcurrentJankenConsumesOnce(t *testing.T) {
	StartJanken(2, 10, "ぱー")
	var wg sync.WaitGroup
	var consumed atomic.Int32
	for range 100 {
		wg.Go(func() {
			if _, ok := PlayJanken(2, 10, true); ok {
				consumed.Add(1)
			}
		})
	}
	wg.Wait()
	if consumed.Load() != 1 {
		t.Fatalf("consumed %d times", consumed.Load())
	}
}
