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

func TestScaleAndOffset(t *testing.T) {
	x := scaleAndOffset(1, 1)
	if x != 1.5 {
		t.Errorf("x!=1 %f", x)
	}
	x = scaleAndOffset(10, 10)
	if x != 105 {
		t.Errorf("x!=105 %f", x)
	}
}
