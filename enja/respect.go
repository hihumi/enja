package enja

import (
	"fmt"
)

type RespectStruct struct {
	respectEnString interface{}
	respectJaString interface{}
	respect         interface{}
}

func RespectEnString() interface{} {
	var RespectEn RespectStruct
	RespectEn.respectEnString = "respect"

	return RespectEn.respectEnString
}

func RespectJaString() interface{} {
	var RespectJa RespectStruct
	RespectJa.respectJaString = `[名]: [C, U]: [主な意味]: 尊重・尊敬
[動]: [主な意味]: ...を尊重・尊敬する`

	return RespectJa.respectJaString

}

func RespectPrintJa() {
	var Respect RespectStruct
	Respect.respect = RespectJaString()

	fmt.Println(Respect.respect)

	return
}
