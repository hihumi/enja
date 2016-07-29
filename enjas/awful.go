package enjas

import (
	"fmt"
)

type AwfulStruct struct {
	awfulEnString, awfulJaString, awful string
}

func AwfulEnString() string {
	var AwfulEn AwfulStruct
	AwfulEn.awfulEnString = "awful"

	return AwfulEn.awfulEnString
}

func AwfulJaString() string {
	var AwfulJa AwfulStruct
	AwfulJa.awfulJaString = `[形]: [主な意味]: 1. ひどい 2. (くだけた表現として)ものすごい`

	return AwfulJa.awfulJaString
}

func AwfulPrintJa() {
	var Awful AwfulStruct
	Awful.awful = AwfulJaString()

	fmt.Println(Awful.awful)

	return
}
