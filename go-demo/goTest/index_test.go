package main

import (
	"testing"
)

func TestHelloWlc(t *testing.T) {
	output := outputWlc()
	expectOutPut := "Wlc"
	// expectOutPut1 := "Tom"

	if output != expectOutPut {
		t.Errorf("Expected do not match")
	}
}