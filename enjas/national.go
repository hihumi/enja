package enjas

import (
	"fmt"
)

type NationalStruct struct {
	nationalEnString, nationalJaString, national string
}

func NationalEnString() string {
	var NationalEn NationalStruct
	NationalEn.nationalEnString = "national"

	return NationalEn.nationalEnString
}

func NationalJaString() string {
	var NationalJa NationalStruct
	NationalJa.nationalJaString = `[形]: [主な意味]: 1. 国の、 国内の、全国の、国家全体の、中央の、2. 国立の、国有の、国営の`

	return NationalJa.nationalJaString
}

func NationalPrintJa() {
	var National NationalStruct
	National.national = NationalJaString()

	fmt.Println(National.national)

	return
}
