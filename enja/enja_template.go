package enja

import (
	"fmt"
)

type EnWordStruct struct {
	enwordEnString, enwordJaString, enword string
}

func EnWordEnString() string {
	var EnWordEn EnWordStruct
	EnWordEn.enwordEnString = "en_word"

	return EnWordEn.enwordEnString
}

func EnWordJaString() string {
	var EnWordJa EnWordStruct
	EnWordJa.enwordJaString = `日本語解説`

	return EnWordJa.enwordJaString
}

func EnWordPrintJa() {
	var EnWord EnWordStruct
	EnWord.enword = EnWordJaString()

	fmt.Println(EnWord.enword)

	return
}
