package main

import (
	"testing"
)

func TestNorm(t *testing.T) {
	got := norm(1, 10)
	if got != 1 {
		t.Fail()
	}

	got = norm(11, 10)
	if got != 1 {
		t.Fail()
	}

	got = norm(-1, 10)
	if got != 9 {
		t.Fail()
	}

	got = norm(-11, 10)
	if got != 9 {
		t.Fail()
	}

	for i := -10; i < 10; i++ {
		got = norm(i, 5)
		if got < 0 || got > 5 {
			t.Fail()
		}
	}
}
