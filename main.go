package main

import (
	"fmt"
)

func main() {
	fmt.Println("enja")
	fmt.Println("英単語の簡素な意味を調べることができます。英単語を小文字で入力してください。終了するには、q、またはQキーを入力してください。")

	for {
		fmt.Printf(">>> ")
		var enWord string
		fmt.Scanln(&enWord)

		q := "q"
		Q := "Q"
		if q == enWord || Q == enWord {
			fmt.Println("end")
			break
		}

		SwitchEnWord(enWord)
	}
}
