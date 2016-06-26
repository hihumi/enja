package enja

import (
	"fmt"
)

type PositiveStruct struct {
	positiveEnString, positiveJaString, positive string
}

func PositiveEnString() string {
	var PositiveEn PositiveStruct
	PositiveEn.positiveEnString = "positive"

	return PositiveEn.positiveEnString
}

func PositiveJaString() string {
	var PositiveJa PositiveStruct
	PositiveJa.positiveJaString = `[形]: [主な意味]: 1. 積極的な、前向きな(の) 2. 確信のある、自信のある 3. 明確な、はっきりした 4. 正(プラス)の性質`

	return PositiveJa.positiveJaString
}

func PositivePrintJa() {
	var Positive PositiveStruct
	Positive.positive = PositiveJaString()

	fmt.Println(Positive.positive)

	return
}
