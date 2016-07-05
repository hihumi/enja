package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("--- Welcome! --- \n\n")
	fmt.Printf("英単語の簡素な意味を調べることができます。英単語を小文字で入力してください。終了するには、q、またはQキーを入力してください。\n")
	fmt.Printf(">>> ")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		enWord := scanner.Text()

		q := "q"
		Q := "Q"
		if q == enWord || Q == enWord {
			fmt.Printf("\n--- See you! --- \n")
			break
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "おっと！ おそらく読み込みエラーです:", err)
		}

		SwitchEnWord(enWord)

		fmt.Printf(">>> ")
	}
}
