package main

import (
	"./enja"
	"fmt"
	_ "github.com/hihumi/enja/enja"
)

func SwitchEnWord(enWord string) {

	var (
		my         = enja.MyEnString()
		respect    = enja.RespectEnString()
		will       = enja.WillEnString()
		individual = enja.IndividualEnString()
		positive   = enja.PositiveEnString()
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
	default:
		fmt.Println("not found")
	}

	return
}
