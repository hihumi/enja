package main

import (
	"./enja"
	"fmt"
)

func SwitchEnWord(enWord string) {

	var (
		my = enja.MyEnString()
	)

	switch enWord {
	case my:
		enja.MyPrintJa()
	default:
		fmt.Println("not found")
	}

	return
}
