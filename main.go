package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("--- Welcome! --- \n\n")

	fmt.Println("英単語の簡素な意味を調べることができます。英単語を小文字で入力してください。")
	fmt.Println("登録されている全英単語を表示したい場合は、words-list と入力してください。")
	fmt.Printf("終了する場合は、q、またはQキーを入力してください。\n\n")

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

		fmt.Printf("\n>>> ")
	}
}
