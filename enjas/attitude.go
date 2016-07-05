package enjas

import (
	"fmt"
)

type AttitudeStruct struct {
	attitudeEnString, attitudeJaString, attitude string
}

func AttitudeEnString() string {
	var AttitudeEn AttitudeStruct
	AttitudeEn.attitudeEnString = "attitude"

	return AttitudeEn.attitudeEnString
}

func AttitudeJaString() string {
	var AttitudeJa AttitudeStruct
	AttitudeJa.attitudeJaString = `[名]: [C, U]: [主な意味]: 態度、姿勢、考え方`

	return AttitudeJa.attitudeJaString
}

func AttitudePrintJa() {
	var Attitude AttitudeStruct
	Attitude.attitude = AttitudeJaString()

	fmt.Println(Attitude.attitude)

	return
}
