package enjas

import (
	"fmt"
)

type EnWordStruct struct {
	enWordEnString, enWordJaString, enWord string
}

func EnWordEnString() string {
	var EnWordEn EnWordStruct
	EnWordEn.enWordEnString = "enWord"

	return EnWordEn.enWordEnString
}

func EnWordJaString() string {
	var EnWordJa EnWordStruct
	EnWordJa.enWordJaString = `日本語解説`

	return EnWordJa.enWordJaString
}

func EnWordPrintJa() {
	var EnWord EnWordStruct
	EnWord.enWord = EnWordJaString()

	fmt.Println(EnWord.enWord)

	return
}
