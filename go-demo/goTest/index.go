package gotest_test

import (
	"testing"
)

func outputWlc() string {
	return "Wlc"
}

func TestHelloWlc(t *testing.T) {
	output := outputWlc()
	expectOutPut := "Tom"

	if output != expectOutPut {
		t.Errorf("Expected do not match")
	}
}