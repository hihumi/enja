package enjas

import (
	"fmt"
)

type LibertyStruct struct {
	libertyEnString, libertyJaString, liberty string
}

func LibertyEnString() string {
	var LibertyEn LibertyStruct
	LibertyEn.libertyEnString = "liberty"

	return LibertyEn.libertyEnString
}

func LibertyJaString() string {
	var LibertyJa LibertyStruct
	LibertyJa.libertyJaString = `[名]: [C, U]: [主な意味]: 自由、独立、権利、特権、特典`

	return LibertyJa.libertyJaString
}

func LibertyPrintJa() {
	var Liberty LibertyStruct
	Liberty.liberty = LibertyJaString()

	fmt.Println(Liberty.liberty)

	return
}
