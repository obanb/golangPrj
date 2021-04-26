package common

import (
	"testing"
)

func Test_something(t *testing.T) {
	val := 1
	val++
	// Assert
	if val == 2 {
		t.Error("failed while testing the new account validation")
	}
}
