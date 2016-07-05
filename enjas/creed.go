package enjas

import (
	"fmt"
)

type CreedStruct struct {
	creedEnString, creedJaString, creed string
}

func CreedEnString() string {
	var CreedEn CreedStruct
	CreedEn.creedEnString = "creed"

	return CreedEn.creedEnString
}

func CreedJaString() string {
	var CreedJa CreedStruct
	CreedJa.creedJaString = `[名]: [C]: [主な意味]: 信条、信仰`

	return CreedJa.creedJaString
}

func CreedPrintJa() {
	var Creed CreedStruct
	Creed.creed = CreedJaString()

	fmt.Println(Creed.creed)

	return
}
