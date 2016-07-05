package enjas

import (
	"fmt"
)

type MyStruct struct {
	myEnString, myJaString, my string
}

func MyEnString() string {
	var MyEn MyStruct
	MyEn.myEnString = "my"

	return MyEn.myEnString
}

func MyJaString() string {
	var MyJa MyStruct
	MyJa.myJaString = "私の"

	return MyJa.myJaString
}

func MyPrintJa() {
	var My MyStruct
	My.my = MyJaString()

	fmt.Printf("[代名詞]: [主な意味]: %v \n(また、[間投詞]としても用いられる)\n", My.my)

	return
}
