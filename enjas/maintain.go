package enjas

import (
	"fmt"
)

type MaintainStruct struct {
	maintainEnString, maintainJaString, maintain string
}

func MaintainEnString() string {
	var MaintainEn MaintainStruct
	MaintainEn.maintainEnString = "maintain"

	return MaintainEn.maintainEnString
}

func MaintainJaString() string {
	var MaintainJa MaintainStruct
	MaintainJa.maintainJaString = `[動]: [主な意味]: 1. (状態や関係など)を維持する、保つ 2. (建物や機械など)を手入れする、整備する 3. 主張する`

	return MaintainJa.maintainJaString
}

func MaintainPrintJa() {
	var Maintain MaintainStruct
	Maintain.maintain = MaintainJaString()

	fmt.Println(Maintain.maintain)

	return
}
