package main

import (
	"fmt"

	"./enjas"
	//_ "github.com/hihumi/enja/enjas"
)

func SwitchEnWord(enWord string) {

	var (
		wordsList = enjas.WordsListEnString()

		// a b c d e f g h i j k l m n o p q r s t u v w x y z
		// a
		attitude = enjas.AttitudeEnString()

		// b

		// c
		creed = enjas.CreedEnString()

		// d

		// e
		equality = enjas.EqualityEnString()

		// f
		fair = enjas.FairEnString()

		// g
		gender = enjas.GenderEnString()

		// h

		// i
		individual = enjas.IndividualEnString()

		// j

		// k

		// l

		// m
		maintain = enjas.MaintainEnString()
		my       = enjas.MyEnString()

		// n
		national = enjas.NationalEnString()
		negative = enjas.NegativeEnString()

		// o
		origin  = enjas.OriginEnString()
		outlook = enjas.OutlookEnString()

		// p
		positive = enjas.PositiveEnString()

		// q

		// r
		respect = enjas.RespectEnString()

		// s

		// t

		// u

		// v

		// w
		will = enjas.WillEnString()

		// x

		// y

		// z
	)

	switch enWord {
	case wordsList:
		enjas.WordsListPrintJa()

	// a b c d e f g h i j k l m n o p q r s t u v w x y z
	// a
	case attitude:
		enjas.AttitudePrintJa()

	// b

	// c
	case creed:
		enjas.CreedPrintJa()

	// d

	// e
	case equality:
		enjas.EqualityPrintJa()

	// f
	case fair:
		enjas.FairPrintJa()

	// g
	case gender:
		enjas.GenderPrintJa()

	// h

	// i
	case individual:
		enjas.IndividualPrintJa()

	// j

	// k

	// l

	// m
	case maintain:
		enjas.MaintainPrintJa()
	case my:
		enjas.MyPrintJa()

	// n
	case national:
		enjas.NationalPrintJa()
	case negative:
		enjas.NegativePrintJa()

	// o
	case origin:
		enjas.OriginPrintJa()
	case outlook:
		enjas.OutlookPrintJa()

	// p
	case positive:
		enjas.PositivePrintJa()

	// q

	// r
	case respect:
		enjas.RespectPrintJa()
	// s

	// t

	// u

	// v

	// w
	case will:
		enjas.WillPrintJa()

	// x

	// y

	// z
	default:
		fmt.Println("not found")
	}

	return
}
