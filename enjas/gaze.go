package enjas

import (
	"fmt"
)

type GazeStruct struct {
	gazeEnString, gazeJaString, gaze string
}

func GazeEnString() string {
	var GazeEn GazeStruct
	GazeEn.gazeEnString = "gaze"

	return GazeEn.gazeEnString
}

func GazeJaString() string {
	var GazeJa GazeStruct
	GazeJa.gazeJaString = `[動]: [主な意味]: じっと見る、見つめる、凝視する、注視する
([名]: 見つめること、凝視、注視)`

	return GazeJa.gazeJaString
}

func GazePrintJa() {
	var Gaze GazeStruct
	Gaze.gaze = GazeJaString()

	fmt.Println(Gaze.gaze)

	return
}
