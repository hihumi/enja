package main

import (
	"fmt"

	"./enja"
	_ "github.com/hihumi/enja/enja"
)

func SwitchEnWord(enWord string) {

	var (
		wordsList = enja.WordsListEnString()

		// a b c d e f g h i j k l m n o p q r s t u v w x y z
		// a
		attitude = enja.AttitudeEnString()

		// b

		// c
		creed = enja.CreedEnString()

		// d

		// e
		equality = enja.EqualityEnString()

		// f
		fair = enja.FairEnString()

		// g
		gender = enja.GenderEnString()

		// h

		// i
		individual = enja.IndividualEnString()

		// j

		// k

		// l

		// m
		maintain = enja.MaintainEnString()
		my       = enja.MyEnString()

		// n
		national = enja.NationalEnString()
		negative = enja.NegativeEnString()

		// o
		origin  = enja.OriginEnString()
		outlook = enja.OutlookEnString()

		// p
		positive = enja.PositiveEnString()

		// q

		// r
		respect = enja.RespectEnString()

		// s

		// t

		// u

		// v

		// w
		will = enja.WillEnString()

		// x

		// y

		// z
	)

	switch enWord {
	case wordsList:
		enja.WordsListPrintJa()

	// a b c d e f g h i j k l m n o p q r s t u v w x y z
	// a
	case attitude:
		enja.AttitudePrintJa()

	// b

	// c
	case creed:
		enja.CreedPrintJa()

	// d

	// e
	case equality:
		enja.EqualityPrintJa()

	// f
	case fair:
		enja.FairPrintJa()

	// g
	case gender:
		enja.GenderPrintJa()

	// h

	// i
	case individual:
		enja.IndividualPrintJa()

	// j

	// k

	// l

	// m
	case maintain:
		enja.MaintainPrintJa()
	case my:
		enja.MyPrintJa()

	// n
	case national:
		enja.NationalPrintJa()
	case negative:
		enja.NegativePrintJa()

	// o
	case origin:
		enja.OriginPrintJa()
	case outlook:
		enja.OutlookPrintJa()

	// p
	case positive:
		enja.PositivePrintJa()

	// q

	// r
	case respect:
		enja.RespectPrintJa()
	// s

	// t

	// u

	// v

	// w
	case will:
		enja.WillPrintJa()

	// x

	// y

	// z
	default:
		fmt.Println("not found")
	}

	return
}
