package handler

import "testing"

func TestIsNumeric(t *testing.T) {
	if got := isNumeric("1"); got != nil {
		t.Errorf("isNumeric(\"1\"): got %v want nil", got)
	}
}
