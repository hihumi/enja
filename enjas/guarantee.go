package enjas

import (
	"fmt"
)

type GuaranteeStruct struct {
	guaranteeEnString, guaranteeJaString, guarantee string
}

func GuaranteeEnString() string {
	var GuaranteeEn GuaranteeStruct
	GuaranteeEn.guaranteeEnString = "guarantee"

	return GuaranteeEn.guaranteeEnString
}

func GuaranteeJaString() string {
	var GuaranteeJa GuaranteeStruct
	GuaranteeJa.guaranteeJaString = `[動]: [主な意味]: ...を保証する、...を約束する、...を保護する`

	return GuaranteeJa.guaranteeJaString
}

func GuaranteePrintJa() {
	var Guarantee GuaranteeStruct
	Guarantee.guarantee = GuaranteeJaString()

	fmt.Println(Guarantee.guarantee)

	return
}
