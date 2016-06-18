package main

import (
	"./enja"
	_ "github.com/hihumi/enja/enja"
	"testing"
)

func TestSwitchEnWord(t *testing.T) {
	// enja.MyEnstring() (my.go: MyEnString())
	myEnExpect := "my"
	myEnActual := enja.MyEnString()

	if myEnExpect != myEnActual {
		t.Errorf("%v != %v\n", myEnExpect, myEnActual)
	}

	// enja.RespectEnString() (respect.go: RespectEnString())
	respectEnExpect := "respect"
	respectEnActual := enja.RespectEnString()

	if respectEnExpect != respectEnActual {
		t.Errorf("%v != %v\n", respectEnExpect, respectEnActual)
	}

	// enja.WillEnString() (will.go: WillEnString())
	willEnExpect := "will"
	willEnActual := enja.WillEnString()

	if willEnExpect != willEnActual {
		t.Errorf("%v != %v\n", willEnExpect, willEnActual)
	}
}
