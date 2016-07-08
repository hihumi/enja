package enjas

import (
	"fmt"
)

type PillarStruct struct {
	pillarEnString, pillarJaString, pillar string
}

func PillarEnString() string {
	var PillarEn PillarStruct
	PillarEn.pillarEnString = "pillar"

	return PillarEn.pillarEnString
}

func PillarJaString() string {
	var PillarJa PillarStruct
	PillarJa.pillarJaString = `[名]: [C]: [主な意味]: 1. (建築物などの)柱、支柱 2. (制度などの)中心部分、要所 3. (人物などの)大黒柱、重要人物、中心人物`

	return PillarJa.pillarJaString
}

func PillarPrintJa() {
	var Pillar PillarStruct
	Pillar.pillar = PillarJaString()

	fmt.Println(Pillar.pillar)

	return
}
