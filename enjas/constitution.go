package enjas

import (
	"fmt"
)

type ConstitutionStruct struct {
	constitutionEnString, constitutionJaString, constitution string
}

func ConstitutionEnString() string {
	var ConstitutionEn ConstitutionStruct
	ConstitutionEn.constitutionEnString = "constitution"

	return ConstitutionEn.constitutionEnString
}

func ConstitutionJaString() string {
	var ConstitutionJa ConstitutionStruct
	ConstitutionJa.constitutionJaString = `[名]: [C]: [主な意味]: 憲法`

	return ConstitutionJa.constitutionJaString
}

func ConstitutionPrintJa() {
	var Constitution ConstitutionStruct
	Constitution.constitution = ConstitutionJaString()

	fmt.Println(Constitution.constitution)

	return
}
