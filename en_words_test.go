package main

import (
	"testing"

	"./enja"
	_ "github.com/hihumi/enja/enja"
)

// template
// enja.EnWordEnstring() (enWord.go: EnWordEnString())
/*
	enWordEnExpect := "enWord"
	enWordEnActual := enja.EnWordEnString()

	if enWordEnExpect != enWordEnActual {
		t.Errorf("%v != %v\n", enWordEnExpect, enWordEnActual)
	}
*/

func TestSwitchEnWord(t *testing.T) {
	// enja.WordListsEnstring() (word_lists.go: WordListsEnString())
	wordsListEnExpect := "words-list"
	wordsListEnActual := enja.WordsListEnString()

	if wordsListEnExpect != wordsListEnActual {
		t.Errorf("%v != %v\n", wordsListEnExpect, wordsListEnActual)
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

	// enja.FairEnstring() (fair.go: FairEnString())
	fairEnExpect := "fair"
	fairEnActual := enja.FairEnString()

	if fairEnExpect != fairEnActual {
		t.Errorf("%v != %v\n", fairEnExpect, fairEnActual)
	}

	// enja.NationalEnstring() (national.go: NationalEnString())
	nationalEnExpect := "national"
	nationalEnActual := enja.NationalEnString()

	if nationalEnExpect != nationalEnActual {
		t.Errorf("%v != %v\n", nationalEnExpect, nationalEnActual)
	}

	// enja.OriginEnstring() (origin.go: OriginEnString())
	originEnExpect := "origin"
	originEnActual := enja.OriginEnString()

	if originEnExpect != originEnActual {
		t.Errorf("%v != %v\n", originEnExpect, originEnActual)
	}

	// enja.GenderEnstring() (gender.go: GenderEnString())
	genderEnExpect := "gender"
	genderEnActual := enja.GenderEnString()

	if genderEnExpect != genderEnActual {
		t.Errorf("%v != %v\n", genderEnExpect, genderEnActual)
	}

	// enja.CreedEnstring() (creed.go: CreedEnString())
	creedEnExpect := "creed"
	creedEnActual := enja.CreedEnString()

	if creedEnExpect != creedEnActual {
		t.Errorf("%v != %v\n", creedEnExpect, creedEnActual)
	}

	//enja.EqualityEnstring() (equality.go: EqualityEnString())
	equalityEnExpect := "equality"
	equalityEnActual := enja.EqualityEnString()

	if equalityEnExpect != equalityEnActual {
		t.Errorf("%v != %v\n", equalityEnExpect, equalityEnActual)
	}

}
