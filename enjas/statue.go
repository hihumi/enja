package enjas

import (
	"fmt"
)

type StatueStruct struct {
	statueEnString, statueJaString, statue string
}

func StatueEnString() string {
	var StatueEn StatueStruct
	StatueEn.statueEnString = "statue"

	return StatueEn.statueEnString
}

func StatueJaString() string {
	var StatueJa StatueStruct
	StatueJa.statueJaString = `[名]: [C]: [主な意味]: (人などの)像、彫像`

	return StatueJa.statueJaString
}

func StatuePrintJa() {
	var Statue StatueStruct
	Statue.statue = StatueJaString()

	fmt.Println(Statue.statue)

	return
}
