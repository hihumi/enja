package main

import (
	"./enja"
	"testing"
)

func TestSwitchEnWord(t *testing.T) {
	// enja.MyEnstring() (my.go: MyEnString())
	MyEnExcept := "my"
	MyEnActual := enja.MyEnString()

	if MyEnExcept != MyEnActual {
		t.Errorf("%v != %v\n", MyEnExcept, MyEnActual)
	}
}
