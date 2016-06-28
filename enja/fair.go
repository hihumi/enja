package enja

import (
	"fmt"
)

type FairStruct struct {
	fairEnString, fairJaString, fair string
}

func FairEnString() string {
	var FairEn FairStruct
	FairEn.fairEnString = "fair"

	return FairEn.fairEnString
}

func FairJaString() string {
	var FairJa FairStruct
	FairJa.fairJaString = `[形]: [主な意味]: 1. 公平な、公正な 2. かなりの 3. まずまずの 4. (天候)晴れた、好天の
(また、[名]: [C]: [主な意味]: 博覧会、見本市、遊園地、説明会、慈善市など)`

	return FairJa.fairJaString
}

func FairPrintJa() {
	var Fair FairStruct
	Fair.fair = FairJaString()

	fmt.Println(Fair.fair)

	return
}
