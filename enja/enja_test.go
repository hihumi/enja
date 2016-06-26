package enja

import (
	"testing"
)

// template
// foo.go
/*
func TestEnwordEnString(t *testing.T) {
	expect := "en_word"
	actual := EnwordEnString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestEnwordJaString(t *testing.T) {
	expect := `日本語解説`
	actual := EnwordJaString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestEnWordPrintJa(t *testing.T) {
	expect := EnwordJaString()
	actual := `日本語解説`

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
*/

// my.go
func TestMyEnString(t *testing.T) {
	expect := "my"
	actual := MyEnString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestMyJaString(t *testing.T) {
	expect := "私の"
	actual := MyJaString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}

// respect.go
func TestRespectEnString(t *testing.T) {
	expect := "respect"
	actual := RespectEnString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestRespectJaString(t *testing.T) {
	expect := `[名]: [C, U]: [主な意味]: 尊重・尊敬
[動]: [主な意味]: ...を尊重・尊敬する`
	actual := RespectJaString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestRespectPrintJa(t *testing.T) {
	expect := RespectJaString()
	actual := `[名]: [C, U]: [主な意味]: 尊重・尊敬
[動]: [主な意味]: ...を尊重・尊敬する`

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}

// will.go
func TestWillEnString(t *testing.T) {
	expect := "will"
	actual := WillEnString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestWillJaString(t *testing.T) {
	expect := `[名]: [C, U]: [主な意味]: 意志
また、[C]: 遺書`
	actual := WillJaString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestWillPrintJa(t *testing.T) {
	expect := WillJaString()
	actual := `[名]: [C, U]: [主な意味]: 意志
また、[C]: 遺書`

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}

// individual.go
func TestIndividualEnString(t *testing.T) {
	expect := "individual"
	actual := IndividualEnString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestIndividualJaString(t *testing.T) {
	expect := `[名]: [C]: [主な意味]: 個人
[形]: [主な意味]: 1.個人の... 2.個々の... 3.個性的な...`
	actual := IndividualJaString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestIndividualPrintJa(t *testing.T) {
	expect := IndividualJaString()
	actual := `[名]: [C]: [主な意味]: 個人
[形]: [主な意味]: 1.個人の... 2.個々の... 3.個性的な...`

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}

// positive.go
func TestPositiveEnString(t *testing.T) {
	expect := "positive"
	actual := PositiveEnString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestPositiveJaString(t *testing.T) {
	expect := `[形]: [主な意味]: 1. 積極的な、前向きな(の) 2. 確信のある、自信のある 3. 明確な、はっきりした 4. 正(プラス)の性質`
	actual := PositiveJaString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestPositivePrintJa(t *testing.T) {
	expect := PositiveJaString()
	actual := `[形]: [主な意味]: 1. 積極的な、前向きな(の) 2. 確信のある、自信のある 3. 明確な、はっきりした 4. 正(プラス)の性質`

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
