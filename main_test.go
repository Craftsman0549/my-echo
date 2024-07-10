package main

import "testing"

func TestPing(t *testing.T) {
	got := ping()
	if got != "hello okayusan" {
		t.Errorf("ping() =%v; want pong", got)
	}
}
