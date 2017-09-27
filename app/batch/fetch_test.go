package main

import "testing"

func TestParseDatetime(t *testing.T) {
	if got, want := parseDatetime(""), "2006-01-02 15:04:05"; got != want {
		t.Errorf("parseDatetime(\"\"): got %v want %v", got, want)
	}
}
