package main

import (
	"testing"

	"./enjas"
	//_ "github.com/hihumi/enja/enjas"
)

// template
// enjas.EnWordEnstring() (enWord.go: EnWordEnString())
/*
	enWordEnExpect := "enWord"
	enWordEnActual := enjas.EnWordEnString()

	if enWordEnExpect != enWordEnActual {
		t.Errorf("%v != %v\n", enWordEnExpect, enWordEnActual)
	}
*/

func TestSwitchEnWord(t *testing.T) {
	// enjas.WordListsEnstring() (word_lists.go: WordListsEnString())
	wordsListEnExpect := "words-list"
	wordsListEnActual := enjas.WordsListEnString()

	if wordsListEnExpect != wordsListEnActual {
		t.Errorf("%v != %v\n", wordsListEnExpect, wordsListEnActual)
	}

	// enjas.MyEnstring() (my.go: MyEnString())
	myEnExpect := "my"
	myEnActual := enjas.MyEnString()

	if myEnExpect != myEnActual {
		t.Errorf("%v != %v\n", myEnExpect, myEnActual)
	}

	// enjas.RespectEnString() (respect.go: RespectEnString())
	respectEnExpect := "respect"
	respectEnActual := enjas.RespectEnString()

	if respectEnExpect != respectEnActual {
		t.Errorf("%v != %v\n", respectEnExpect, respectEnActual)
	}

	// enjas.WillEnString() (will.go: WillEnString())
	willEnExpect := "will"
	willEnActual := enjas.WillEnString()

	if willEnExpect != willEnActual {
		t.Errorf("%v != %v\n", willEnExpect, willEnActual)
	}

	// enjas.IndividualEnString() (individual.go: IndividualEnString())
	individualEnExpect := "individual"
	individualEnActual := enjas.IndividualEnString()

	if individualEnExpect != individualEnActual {
		t.Errorf("%v != %v\n", individualEnExpect, individualEnActual)
	}

	// enjas.PositiveEnString() (positive.go: PositiveEnString())
	positiveEnExpect := "positive"
	positiveEnActual := enjas.PositiveEnString()

	if positiveEnExpect != positiveEnActual {
		t.Errorf("%v != %v\n", positiveEnExpect, positiveEnActual)
	}

	// enjas.NegativeEnstring() (negative.go: NegativeEnString())
	negativeEnExpect := "negative"
	negativeEnActual := enjas.NegativeEnString()

	if negativeEnExpect != negativeEnActual {
		t.Errorf("%v != %v\n", negativeEnExpect, negativeEnActual)
	}

	// enjas.OutlookEnstring() (outlook.go: OutlookEnString())
	outlookEnExpect := "outlook"
	outlookEnActual := enjas.OutlookEnString()

	if outlookEnExpect != outlookEnActual {
		t.Errorf("%v != %v\n", outlookEnExpect, outlookEnActual)
	}

	// enjas.MaintainEnstring() (maintain.go: MaintainEnString())
	maintainEnExpect := "maintain"
	maintainEnActual := enjas.MaintainEnString()

	if maintainEnExpect != maintainEnActual {
		t.Errorf("%v != %v\n", maintainEnExpect, maintainEnActual)
	}

	// enjas.AttitudeEnstring() (attitude.go: AttitudeEnString())
	attitudeEnExpect := "attitude"
	attitudeEnActual := enjas.AttitudeEnString()

	if attitudeEnExpect != attitudeEnActual {
		t.Errorf("%v != %v\n", attitudeEnExpect, attitudeEnActual)
	}

	// enjas.FairEnstring() (fair.go: FairEnString())
	fairEnExpect := "fair"
	fairEnActual := enjas.FairEnString()

	if fairEnExpect != fairEnActual {
		t.Errorf("%v != %v\n", fairEnExpect, fairEnActual)
	}

	// enjas.NationalEnstring() (national.go: NationalEnString())
	nationalEnExpect := "national"
	nationalEnActual := enjas.NationalEnString()

	if nationalEnExpect != nationalEnActual {
		t.Errorf("%v != %v\n", nationalEnExpect, nationalEnActual)
	}

	// enjas.OriginEnstring() (origin.go: OriginEnString())
	originEnExpect := "origin"
	originEnActual := enjas.OriginEnString()

	if originEnExpect != originEnActual {
		t.Errorf("%v != %v\n", originEnExpect, originEnActual)
	}

	// enjas.GenderEnstring() (gender.go: GenderEnString())
	genderEnExpect := "gender"
	genderEnActual := enjas.GenderEnString()

	if genderEnExpect != genderEnActual {
		t.Errorf("%v != %v\n", genderEnExpect, genderEnActual)
	}

	// enjas.CreedEnstring() (creed.go: CreedEnString())
	creedEnExpect := "creed"
	creedEnActual := enjas.CreedEnString()

	if creedEnExpect != creedEnActual {
		t.Errorf("%v != %v\n", creedEnExpect, creedEnActual)
	}

	//enjas.EqualityEnstring() (equality.go: EqualityEnString())
	equalityEnExpect := "equality"
	equalityEnActual := enjas.EqualityEnString()

	if equalityEnExpect != equalityEnActual {
		t.Errorf("%v != %v\n", equalityEnExpect, equalityEnActual)
	}

}
