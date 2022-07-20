package utils

import "testing"

func TestIsEmpty(t *testing.T) {
	if !IsEmpty("  ") {
		t.Fail()
	}

	if IsEmpty(" nonmpty ") {
		t.Fail()
	}
}
