package enjas

import (
	"fmt"
)

type PerfumeStruct struct {
	perfumeEnString, perfumeJaString, perfume string
}

func PerfumeEnString() string {
	var PerfumeEn PerfumeStruct
	PerfumeEn.perfumeEnString = "perfume"

	return PerfumeEn.perfumeEnString
}

func PerfumeJaString() string {
	var PerfumeJa PerfumeStruct
	PerfumeJa.perfumeJaString = `[名]: [C, U]: [主な意味]: 1. 香水、香料、芳香剤 2. 香り、芳香`

	return PerfumeJa.perfumeJaString
}

func PerfumePrintJa() {
	var Perfume PerfumeStruct
	Perfume.perfume = PerfumeJaString()

	fmt.Println(Perfume.perfume)

	return
}
