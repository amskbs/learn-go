package main

import (
	"testing"
)

func TestFooFunc(t *testing.T) {
	expectedFooResult := "hi"
	if actualFooResult := test(); actualFooResult != expectedFooResult {
		t.Errorf("expected %s; got: %s", expectedFooResult, actualFooResult)
	}
}
