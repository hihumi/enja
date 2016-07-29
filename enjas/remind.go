package enjas

import (
	"fmt"
)

type RemindStruct struct {
	remindEnString, remindJaString, remind string
}

func RemindEnString() string {
	var RemindEn RemindStruct
	RemindEn.remindEnString = "remind"

	return RemindEn.remindEnString
}

func RemindJaString() string {
	var RemindJa RemindStruct
	RemindJa.remindJaString = `[動]: [主な意味]: 1. ...に思い出させる 2. ...に気付かせる`

	return RemindJa.remindJaString
}

func RemindPrintJa() {
	var Remind RemindStruct
	Remind.remind = RemindJaString()

	fmt.Println(Remind.remind)

	return
}
