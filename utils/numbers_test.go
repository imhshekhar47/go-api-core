package utils

import "testing"

func TestGetNumOrElse(t *testing.T) {

	if GetNumOrElse(1, 5) != 1 {
		t.Fail()
	}

	if GetNumOrElse(0, 5) != 5 {
		t.Fail()
	}

	if GetNumOrElse(1.0, 5.0) != 1.0 {
		t.Fail()
	}

	if GetNumOrElse(0.0, 5.0) != 5.0 {
		t.Fail()
	}
}
