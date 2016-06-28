package main

import (
	"testing"

	"./enja"
	_ "github.com/hihumi/enja/enja"
)

// template
// enja.FooEnstring() (foo.go: FooEnString())
/*
	fooEnExpect := "en_word"
	fooEnActual := enja.FooEnString()

	if fooEnExpect != fooEnActual {
		t.Errorf("%v != %v\n", fooEnExpect, fooEnActual)
	}
*/

func TestSwitchEnWord(t *testing.T) {
	// enja.WordListsEnstring() (word_lists.go: WordListsEnString())
	wordListsEnExpect := "word-lists"
	wordListsEnActual := enja.WordListsEnString()

	if wordListsEnExpect != wordListsEnActual {
		t.Errorf("%v != %v\n", wordListsEnExpect, wordListsEnActual)
	}

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

	// enja.IndividualEnString() (individual.go: IndividualEnString())
	individualEnExpect := "individual"
	individualEnActual := enja.IndividualEnString()

	if individualEnExpect != individualEnActual {
		t.Errorf("%v != %v\n", individualEnExpect, individualEnActual)
	}

	// enja.PositiveEnString() (positive.go: PositiveEnString())
	positiveEnExpect := "positive"
	positiveEnActual := enja.PositiveEnString()

	if positiveEnExpect != positiveEnActual {
		t.Errorf("%v != %v\n", positiveEnExpect, positiveEnActual)
	}

	// enja.NegativeEnstring() (negative.go: NegativeEnString())
	negativeEnExpect := "negative"
	negativeEnActual := enja.NegativeEnString()

	if negativeEnExpect != negativeEnActual {
		t.Errorf("%v != %v\n", negativeEnExpect, negativeEnActual)
	}

	// enja.OutlookEnstring() (outlook.go: OutlookEnString())
	outlookEnExpect := "outlook"
	outlookEnActual := enja.OutlookEnString()

	if outlookEnExpect != outlookEnActual {
		t.Errorf("%v != %v\n", outlookEnExpect, outlookEnActual)
	}

	// enja.MaintainEnstring() (maintain.go: MaintainEnString())
	maintainEnExpect := "maintain"
	maintainEnActual := enja.MaintainEnString()

	if maintainEnExpect != maintainEnActual {
		t.Errorf("%v != %v\n", maintainEnExpect, maintainEnActual)
	}

    // enja.AttitudeEnstring() (attitude.go: AttitudeEnString())
	attitudeEnExpect := "attitude"
	attitudeEnActual := enja.AttitudeEnString()

	if attitudeEnExpect != attitudeEnActual {
		t.Errorf("%v != %v\n", attitudeEnExpect, attitudeEnActual)
	}
}
