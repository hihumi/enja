package enjas

import (
	"fmt"
)

type GenderStruct struct {
	genderEnString, genderJaString, gender string
}

func GenderEnString() string {
	var GenderEn GenderStruct
	GenderEn.genderEnString = "gender"

	return GenderEn.genderEnString
}

func GenderJaString() string {
	var GenderJa GenderStruct
	GenderJa.genderJaString = `[名]: [C, U]: [主な意味]: (社会的・文化的に形成される)性別
(なお、sexは、生物的・肉体的な性別)`

	return GenderJa.genderJaString
}

func GenderPrintJa() {
	var Gender GenderStruct
	Gender.gender = GenderJaString()

	fmt.Println(Gender.gender)

	return
}
