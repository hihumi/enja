package enja

import (
	"fmt"
)

type WillStruct struct {
	willEnString, willJaString, will string
}

func WillEnString() string {
	var WillEn WillStruct
	WillEn.willEnString = "will"

	return WillEn.willEnString
}

func WillJaString() string {
	var WillJa WillStruct
	WillJa.willJaString = `[名]: [C, U]: [主な意味]: 意志
また、[C]: 遺書`

	return WillJa.willJaString
}

func WillPrintJa() {
	var Will WillStruct
	Will.will = WillJaString()

	fmt.Println(Will.will)

	return
}
