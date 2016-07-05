package enjas

import (
	"fmt"
)

type NegativeStruct struct {
	negativeEnString, negativeJaString, negative string
}

func NegativeEnString() string {
	var NegativeEn NegativeStruct
	NegativeEn.negativeEnString = "negative"

	return NegativeEn.negativeEnString
}

func NegativeJaString() string {
	var NegativeJa NegativeStruct
	NegativeJa.negativeJaString = `[形]: [主な意味]: 1. 悲観的な、消極的な 2. 否定の 3. 負(マイナス)の性質
(また、写真の「ネガ」: a negative film)`

	return NegativeJa.negativeJaString
}

func NegativePrintJa() {
	var Negative NegativeStruct
	Negative.negative = NegativeJaString()

	fmt.Println(Negative.negative)

	return
}
