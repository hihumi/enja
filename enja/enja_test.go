package enja

import (
	"testing"
)

// my.go
func TestMyEnString(t *testing.T) {
	expect := "my"
	actual := MyEnString()

	if expect != actual {
		t.Errorf("%v != %v\n", expect, actual)
	}
}
func TestMyJaString(t *testing.T) {
	expect := "私の"
	actual := MyJaString()

	if expect != actual {
		t.Errorf("%v != %v\n", expect, actual)
	}
}
