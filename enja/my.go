package enja

import (
	"fmt"
)

type MyStruct struct {
	myEnString interface{}
	myJaString interface{}
	my         interface{}
}

func MyEnString() interface{} {
	var MyEn MyStruct
	MyEn.myEnString = "my"

	return MyEn.myEnString
}

func MyJaString() interface{} {
	var MyJa MyStruct
	MyJa.myJaString = "私の"

	return MyJa.myJaString
}

func MyPrintJa() {
	var My MyStruct
	My.my = MyJaString()

	fmt.Printf("[代名詞]: \n[主な意味]: %v \n(また、[間投詞]としても用いられる)\n", My.my)

	return
}
