package enjas

import (
	"fmt"
)

type OriginStruct struct {
	originEnString, originJaString, origin string
}

func OriginEnString() string {
	var OriginEn OriginStruct
	OriginEn.originEnString = "origin"

	return OriginEn.originEnString
}

func OriginJaString() string {
	var OriginJa OriginStruct
	OriginJa.originJaString = `[名]: [C, U]: 起源、発祥、生まれ`

	return OriginJa.originJaString
}

func OriginPrintJa() {
	var Origin OriginStruct
	Origin.origin = OriginJaString()

	fmt.Println(Origin.origin)

	return
}
