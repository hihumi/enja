package enjas

import (
	"fmt"
)

type SubtleStruct struct {
	subtleEnString, subtleJaString, subtle string
}

func SubtleEnString() string {
	var SubtleEn SubtleStruct
	SubtleEn.subtleEnString = "subtle"

	return SubtleEn.subtleEnString
}

func SubtleJaString() string {
	var SubtleJa SubtleStruct
	SubtleJa.subtleJaString = `[形]: [主な意味]: 1. 微かな、微妙な 2. 巧妙な 3. 鋭敏な`

	return SubtleJa.subtleJaString
}

func SubtlePrintJa() {
	var Subtle SubtleStruct
	Subtle.subtle = SubtleJaString()

	fmt.Println(Subtle.subtle)

	return
}
