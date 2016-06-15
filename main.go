package main

import (
	"fmt"
)

func main() {
	fmt.Println("enja")

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

		const (
			enS1 = "mean"
			enS2 = "my"
		)

		const (
			jaS1 = "平均"
			jaS2 = "私"
		)

		switch enWord {
		case enS1:
			fmt.Println(jaS1)
		case enS2:
			fmt.Println(jaS2)
		default:
			fmt.Println("not found")
		}

	}
}
