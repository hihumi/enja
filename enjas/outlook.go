package enjas

import (
	"fmt"
)

type OutlookStruct struct {
	outlookEnString, outlookJaString, outlook string
}

func OutlookEnString() string {
	var OutlookEn OutlookStruct
	OutlookEn.outlookEnString = "outlook"

	return OutlookEn.outlookEnString
}

func OutlookJaString() string {
	var OutlookJa OutlookStruct
	OutlookJa.outlookJaString = `[名]: [C]: [主な意味]: 1. 見通し 2. 視野 3. 展望`

	return OutlookJa.outlookJaString
}

func OutlookPrintJa() {
	var Outlook OutlookStruct
	Outlook.outlook = OutlookJaString()

	fmt.Println(Outlook.outlook)

	return
}
