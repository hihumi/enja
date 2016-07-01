package main

import (
	"fmt"

	"./enja"
	_ "github.com/hihumi/enja/enja"
)

func SwitchEnWord(enWord string) {

	var (
		wordsList = enja.WordsListEnString()

		// a b c d e f g h i j k l m n o p q r s t u v w x y z
		attitude   = enja.AttitudeEnString()
		fair       = enja.FairEnString()
		gender     = enja.GenderEnString()
		individual = enja.IndividualEnString()
		maintain   = enja.MaintainEnString()
		my         = enja.MyEnString()
		national   = enja.NationalEnString()
		negative   = enja.NegativeEnString()
		origin     = enja.OriginEnString()
		outlook    = enja.OutlookEnString()
		positive   = enja.PositiveEnString()
		respect    = enja.RespectEnString()
		will       = enja.WillEnString()
	)

	switch enWord {
	case wordsList:
		enja.WordsListPrintJa()

	// a b c d e f g h i j k l m n o p q r s t u v w x y z
	case attitude:
		enja.AttitudePrintJa()
	case fair:
		enja.FairPrintJa()
	case gender:
		enja.GenderPrintJa()
	case individual:
		enja.IndividualPrintJa()
	case maintain:
		enja.MaintainPrintJa()
	case my:
		enja.MyPrintJa()
	case national:
		enja.NationalPrintJa()
	case negative:
		enja.NegativePrintJa()
	case origin:
		enja.OriginPrintJa()
	case outlook:
		enja.OutlookPrintJa()
	case positive:
		enja.PositivePrintJa()
	case respect:
		enja.RespectPrintJa()
	case will:
		enja.WillPrintJa()
	default:
		fmt.Println("not found")
	}

	return
}
