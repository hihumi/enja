package main

import (
	"fmt"

	"./enja"
	_ "github.com/hihumi/enja/enja"
)

func SwitchEnWord(enWord string) {

	var (
		wordLists = enja.WordListsEnString()

		attitude   = enja.AttitudeEnString()
		fair       = enja.FairEnString()
		individual = enja.IndividualEnString()
		maintain   = enja.MaintainEnString()
		my         = enja.MyEnString()
		national   = enja.NationalEnString()
		negative   = enja.NegativeEnString()
		outlook    = enja.OutlookEnString()
		positive   = enja.PositiveEnString()
		respect    = enja.RespectEnString()
		will       = enja.WillEnString()
	)

	switch enWord {
	case wordLists:
		enja.WordListsPrintJa()
	case attitude:
		enja.AttitudePrintJa()
	case fair:
		enja.FairPrintJa()
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
