package main

import (
	"fmt"

	"./enja"
	_ "github.com/hihumi/enja/enja"
)

func SwitchEnWord(enWord string) {

	var (
		wordLists = enja.WordListsEnString()

		my         = enja.MyEnString()
		respect    = enja.RespectEnString()
		will       = enja.WillEnString()
		individual = enja.IndividualEnString()
		positive   = enja.PositiveEnString()
		negative   = enja.NegativeEnString()
		outlook    = enja.OutlookEnString()
		maintain   = enja.MaintainEnString()
	)

	switch enWord {
	case wordLists:
		enja.WordListsPrintJa()
	case my:
		enja.MyPrintJa()
	case respect:
		enja.RespectPrintJa()
	case will:
		enja.WillPrintJa()
	case individual:
		enja.IndividualPrintJa()
	case positive:
		enja.PositivePrintJa()
	case negative:
		enja.NegativePrintJa()
	case outlook:
		enja.OutlookPrintJa()
	case maintain:
		enja.MaintainPrintJa()
	default:
		fmt.Println("not found")
	}

	return
}
