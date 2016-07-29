package enjas

import (
	"fmt"
)

type ScentStruct struct {
	scentEnString, scentJaString, scent string
}

func ScentEnString() string {
	var ScentEn ScentStruct
	ScentEn.scentEnString = "scent"

	return ScentEn.scentEnString
}

func ScentJaString() string {
	var ScentJa ScentStruct
	ScentJa.scentJaString = `[名]: [C, U]: [主な意味]: 1. 香り、匂い 2. 臭い、臭跡`

	return ScentJa.scentJaString
}

func ScentPrintJa() {
	var Scent ScentStruct
	Scent.scent = ScentJaString()

	fmt.Println(Scent.scent)

	return
}
