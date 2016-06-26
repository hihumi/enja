package main

import (
    "fmt"

	"./enja"
	_ "github.com/hihumi/enja/enja"
)

func SwitchEnWord(enWord string) {

	var (
		my         = enja.MyEnString()
		respect    = enja.RespectEnString()
		will       = enja.WillEnString()
		individual = enja.IndividualEnString()
		positive   = enja.PositiveEnString()
		negative   = enja.NegativeEnString()
        outlook = enja.OutlookEnString()
	)

	switch enWord {
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
	default:
		fmt.Println("not found")
	}

	return
}
