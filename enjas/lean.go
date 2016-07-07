package enjas

import (
	"fmt"
)

type LeanStruct struct {
	leanEnString, leanJaString, lean string
}

func LeanEnString() string {
	var LeanEn LeanStruct
	LeanEn.leanEnString = "lean"

	return LeanEn.leanEnString
}

func LeanJaString() string {
	var LeanJa LeanStruct
	LeanJa.leanJaString = `[動]: [主な意味]: 1. ...傾ける、2. 傾く 3. 寄りかかる
([形]: 1. (身体などが)引き締まった 2. (身体などが)痩せた 3. (身体などが)細い)`

	return LeanJa.leanJaString
}

func LeanPrintJa() {
	var Lean LeanStruct
	Lean.lean = LeanJaString()

	fmt.Println(Lean.lean)

	return
}
