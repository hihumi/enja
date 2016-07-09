package main

import (
	"fmt"

	"./enjas"
	// _ "github.com/hihumi/enja/enjas"
)

func SwitchEnWord(enWord string) {

	var (
		wordsList = enjas.WordsListEnString()

		// a b c d e f g h i j k l m n o p q r s t u v w x y z
		// a
		attitude = enjas.AttitudeEnString()

		// b

		// c
		constitution = enjas.ConstitutionEnString()
		creed        = enjas.CreedEnString()

		// d

		// e
		equality = enjas.EqualityEnString()

		// f
		fair = enjas.FairEnString()

		// g
		gaze      = enjas.GazeEnString()
		gender    = enjas.GenderEnString()
		guarantee = enjas.GuaranteeEnString()

		// h

		// i
		individual = enjas.IndividualEnString()

		// j

		// k

		// l
		lean    = enjas.LeanEnString()
		liberty = enjas.LibertyEnString()

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
		pillar   = enjas.PillarEnString()
		positive = enjas.PositiveEnString()

		// q

		// r
		respect = enjas.RespectEnString()

		// s
		statue = enjas.StatueEnString()

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
		enjas.WordsListPrintString()

	// a b c d e f g h i j k l m n o p q r s t u v w x y z
	// a
	case attitude:
		enjas.AttitudePrintJa()

	// b

	// c
	case constitution:
		enjas.ConstitutionPrintJa()
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
	case gaze:
		enjas.GazePrintJa()
	case gender:
		enjas.GenderPrintJa()
	case guarantee:
		enjas.GuaranteePrintJa()

	// h

	// i
	case individual:
		enjas.IndividualPrintJa()

	// j

	// k

	// l
	case lean:
		enjas.LeanPrintJa()
	case liberty:
		enjas.LibertyPrintJa()

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
	case pillar:
		enjas.PillarPrintJa()
	case positive:
		enjas.PositivePrintJa()

	// q

	// r
	case respect:
		enjas.RespectPrintJa()

	// s
	case statue:
		enjas.StatuePrintJa()

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
