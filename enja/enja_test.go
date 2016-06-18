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
	expect := `[名]: [C, U]: [主な意味]: 尊重・尊敬
[動]: [主な意味]: ...を尊重・尊敬する`
	actual := RespectJaString()

	if expect != actual {
		t.Errorf("%v != %v\n", expect, actual)
	}
}
func TestRespectPrintJa(t *testing.T) {
	expect := RespectJaString()
	actual := `[名]: [C, U]: [主な意味]: 尊重・尊敬
[動]: [主な意味]: ...を尊重・尊敬する`

	if expect != actual {
		t.Errorf("%v != %v\n", expect, actual)
	}
}

// will.go
func TestWillEnString(t *testing.T) {
	expect := "will"
	actual := WillEnString()

	if expect != actual {
		t.Errorf("%v != %v\n", expect, actual)
	}
}
func TestWillJaString(t *testing.T) {
	expect := `[名]: [C, U]: [主な意味]: 意志
また、[C]: 遺書`
	actual := WillJaString()

	if expect != actual {
		t.Errorf("%v != %v\n", expect, actual)
	}
}
func TestWillPrintJa(t *testing.T) {
	expect := WillJaString()
	actual := `[名]: [C, U]: [主な意味]: 意志
また、[C]: 遺書`

	if expect != actual {
		t.Errorf("%v != %v\n", expect, actual)
	}
}

// individual.go
func TestIndividualEnString(t *testing.T) {
	expect := "individual"
	actual := IndividualEnString()

	if expect != actual {
		t.Errorf("%v != %v\n", expect, actual)
	}
}
func TestIndividualJaString(t *testing.T) {
	expect := `[名]: [C]: [主な意味]: 個人
[形]: [主な意味]: 1.個人の... 2.個々の... 3.個性的な...`
	actual := IndividualJaString()

	if expect != actual {
		t.Errorf("%v != %v\n", expect, actual)
	}
}
func TestIndividualPrintJa(t *testing.T) {
	expect := IndividualJaString()
	actual := `[名]: [C]: [主な意味]: 個人
[形]: [主な意味]: 1.個人の... 2.個々の... 3.個性的な...`

	if expect != actual {
		t.Errorf("%v != %v\n", expect, actual)
	}
}
