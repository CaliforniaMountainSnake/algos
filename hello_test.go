package main

import "testing"

func TestMax(t *testing.T) {
	got := Max(10, 20)
	want := 20

	if got != want {
		t.Errorf("Max(10, 20) = %d; want %d", got, want)
	}
}
