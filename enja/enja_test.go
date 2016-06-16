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

// respect.go
func TestRespectEnString(t *testing.T) {
	expect := "respect"
	actual := RespectEnString()

	if expect != actual {
		t.Errorf("%v != %v\n", expect, actual)
	}
}
func TestRespectJaString(t *testing.T) {
	expect := "尊重・尊敬"
	actual := RespectJaString()

	if expect != actual {
		t.Errorf("%v != %v\n", expect, actual)
	}
}
