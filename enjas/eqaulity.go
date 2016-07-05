package enjas

import (
	"fmt"
)

type EqualityStruct struct {
	equalityEnString, equalityJaString, equality string
}

func EqualityEnString() string {
	var EqualityEn EqualityStruct
	EqualityEn.equalityEnString = "equality"

	return EqualityEn.equalityEnString
}

func EqualityJaString() string {
	var EqualityJa EqualityStruct
	EqualityJa.equalityJaString = `[名]: [U]: [主な意味]: 1. 平等、均等 2. 均質性、一様性 3. (数学)等式`

	return EqualityJa.equalityJaString
}

func EqualityPrintJa() {
	var Equality EqualityStruct
	Equality.equality = EqualityJaString()

	fmt.Println(Equality.equality)

	return
}
