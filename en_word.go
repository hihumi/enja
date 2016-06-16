package main

import (
	"github.com/hihumi/enja"
	"fmt"
)

func SwitchEnWord(enWord string) {

	var (
		my      = enja.MyEnString()
		respect = enja.RespectEnString()
	)

	switch enWord {
	case my:
		enja.MyPrintJa()
	case respect:
		enja.RespectPrintJa()
	default:
		fmt.Println("not found")
	}

	return
}
