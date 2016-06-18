package main

import (
	"./enja"
	"fmt"
	_ "github.com/hihumi/enja/enja"
)

func SwitchEnWord(enWord string) {

	var (
		my      = enja.MyEnString()
		respect = enja.RespectEnString()
		will    = enja.WillEnString()
	)

	switch enWord {
	case my:
		enja.MyPrintJa()
	case respect:
		enja.RespectPrintJa()
	case will:
		enja.WillPrintJa()
	default:
		fmt.Println("not found")
	}

	return
}
