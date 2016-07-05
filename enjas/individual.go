package enjas

import (
	"fmt"
)

type IndividualStruct struct {
	individualEnString, individualJaString, individual string
}

func IndividualEnString() string {
	var IndividualEn IndividualStruct
	IndividualEn.individualEnString = "individual"

	return IndividualEn.individualEnString
}

func IndividualJaString() string {
	var IndividualJa IndividualStruct
	IndividualJa.individualJaString = `[名]: [C]: [主な意味]: 個人
[形]: [主な意味]: 1.個人の... 2.個々の... 3.個性的な...`

	return IndividualJa.individualJaString
}

func IndividualPrintJa() {
	var Individual IndividualStruct
	Individual.individual = IndividualJaString()

	fmt.Println(Individual.individual)

	return
}
