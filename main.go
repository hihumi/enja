package main

import (
	"fmt"
)

func main() {
	fmt.Printf("--- Welcome! --- \n\n")
	fmt.Printf("英単語の簡素な意味を調べることができます。英単語を小文字で入力してください。終了するには、q、またはQキーを入力してください。\n")

	for {
		fmt.Printf(">>> ")
		var enWord string
		fmt.Scanln(&enWord)

		q := "q"
		Q := "Q"
		if q == enWord || Q == enWord {
			fmt.Printf("\n--- See you! --- \n")
			break
		}

		SwitchEnWord(enWord)
	}
}
