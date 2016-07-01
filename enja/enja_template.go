package enja

import (
	"fmt"
)

type EnwordStruct struct {
	enwordEnString, enwordJaString, enword string
}

func EnwordEnString() string {
	var EnwordEn EnwordStruct
	EnwordEn.enwordEnString = "enword"

	return EnwordEn.enwordEnString
}

func EnwordJaString() string {
	var EnwordJa EnwordStruct
	EnwordJa.enwordJaString = `日本語解説`

	return EnwordJa.enwordJaString
}

func EnwordPrintJa() {
	var Enword EnwordStruct
	Enword.enword = EnwordJaString()

	fmt.Println(Enword.enword)

	return
}
