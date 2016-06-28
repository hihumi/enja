package enja

import (
	"testing"
)

// template
// foo.go
/*
func TestEnwordEnString(t *testing.T) {
	expect := "en_word"
	actual := EnwordEnString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestEnwordJaString(t *testing.T) {
	expect := `日本語解説`
	actual := EnwordJaString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestEnWordPrintJa(t *testing.T) {
	expect := EnwordJaString()
	actual := `日本語解説`

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
*/

// word_lists.go
func TestWordListsEnString(t *testing.T) {
	expect := "word-lists"
	actual := WordListsEnString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestWordListsJaString(t *testing.T) {
	expect := `word lists:
[a]:
  attitude,

[b]:

[c]:

[d]:

[e]:

[f]:
  fair,

[g]:

[h]:

[i]:
  individual,

[j]:

[k]:

[l]:

[m]:
  my,
  maintain,

[n]:
  negative,

[o]:
  outlook,

[p]:
  positive,

[q]:

[r]:
  respect,

[s]:

[t]:

[u]:

[v]:

[w]:
  will,

[x]:

[y]:

[z]:`
	actual := WordListsJaString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestWordListsPrintJa(t *testing.T) {
	expect := WordListsJaString()
	actual := `word lists:
[a]:
  attitude,

[b]:

[c]:

[d]:

[e]:

[f]:
  fair,

[g]:

[h]:

[i]:
  individual,

[j]:

[k]:

[l]:

[m]:
  my,
  maintain,

[n]:
  negative,

[o]:
  outlook,

[p]:
  positive,

[q]:

[r]:
  respect,

[s]:

[t]:

[u]:

[v]:

[w]:
  will,

[x]:

[y]:

[z]:`

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}

// my.go
func TestMyEnString(t *testing.T) {
	expect := "my"
	actual := MyEnString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestMyJaString(t *testing.T) {
	expect := "私の"
	actual := MyJaString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}

// respect.go
func TestRespectEnString(t *testing.T) {
	expect := "respect"
	actual := RespectEnString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestRespectJaString(t *testing.T) {
	expect := `[名]: [C, U]: [主な意味]: 尊重・尊敬
[動]: [主な意味]: ...を尊重・尊敬する`
	actual := RespectJaString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestRespectPrintJa(t *testing.T) {
	expect := RespectJaString()
	actual := `[名]: [C, U]: [主な意味]: 尊重・尊敬
[動]: [主な意味]: ...を尊重・尊敬する`

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}

// will.go
func TestWillEnString(t *testing.T) {
	expect := "will"
	actual := WillEnString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestWillJaString(t *testing.T) {
	expect := `[名]: [C, U]: [主な意味]: 意志
また、[C]: 遺書`
	actual := WillJaString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestWillPrintJa(t *testing.T) {
	expect := WillJaString()
	actual := `[名]: [C, U]: [主な意味]: 意志
また、[C]: 遺書`

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}

// individual.go
func TestIndividualEnString(t *testing.T) {
	expect := "individual"
	actual := IndividualEnString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestIndividualJaString(t *testing.T) {
	expect := `[名]: [C]: [主な意味]: 個人
[形]: [主な意味]: 1.個人の... 2.個々の... 3.個性的な...`
	actual := IndividualJaString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestIndividualPrintJa(t *testing.T) {
	expect := IndividualJaString()
	actual := `[名]: [C]: [主な意味]: 個人
[形]: [主な意味]: 1.個人の... 2.個々の... 3.個性的な...`

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}

// positive.go
func TestPositiveEnString(t *testing.T) {
	expect := "positive"
	actual := PositiveEnString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestPositiveJaString(t *testing.T) {
	expect := `[形]: [主な意味]: 1. 積極的な、前向きな(の) 2. 確信のある、自信のある 3. 明確な、はっきりした 4. 正(プラス)の性質`
	actual := PositiveJaString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestPositivePrintJa(t *testing.T) {
	expect := PositiveJaString()
	actual := `[形]: [主な意味]: 1. 積極的な、前向きな(の) 2. 確信のある、自信のある 3. 明確な、はっきりした 4. 正(プラス)の性質`

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}

// negative.go
func TestNegativeEnString(t *testing.T) {
	expect := "negative"
	actual := NegativeEnString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestNegativeJaString(t *testing.T) {
	expect := `[形]: [主な意味]: 1. 悲観的な、消極的な 2. 否定の 3. 負(マイナス)の性質
(また、写真の「ネガ」: a negative film)`

	actual := NegativeJaString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestNegativePrintJa(t *testing.T) {
	expect := NegativeJaString()
	actual := `[形]: [主な意味]: 1. 悲観的な、消極的な 2. 否定の 3. 負(マイナス)の性質
(また、写真の「ネガ」: a negative film)`

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}

// outlook.go
func TestOutlookEnString(t *testing.T) {
	expect := "outlook"
	actual := OutlookEnString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestOutlookJaString(t *testing.T) {
	expect := `[名]: [C]: [主な意味]: 1. 見通し 2. 視野 3. 展望`
	actual := OutlookJaString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestOutlookPrintJa(t *testing.T) {
	expect := OutlookJaString()
	actual := `[名]: [C]: [主な意味]: 1. 見通し 2. 視野 3. 展望`

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}

// maintain.go
func TestMaintainEnString(t *testing.T) {
	expect := "maintain"
	actual := MaintainEnString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestMaintainJaString(t *testing.T) {
	expect := `[動]: [主な意味]: 1. (状態や関係など)を維持する、保つ 2. (建物や機械など)を手入れする、整備する 3. 主張する`
	actual := MaintainJaString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestMaintainPrintJa(t *testing.T) {
	expect := MaintainJaString()
	actual := `[動]: [主な意味]: 1. (状態や関係など)を維持する、保つ 2. (建物や機械など)を手入れする、整備する 3. 主張する`

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}

// attitude.go
func TestAttitudeEnString(t *testing.T) {
	expect := "attitude"
	actual := AttitudeEnString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestAttitudeJaString(t *testing.T) {
	expect := `[名]: [C, U]: [主な意味]: 態度、姿勢、考え方`
	actual := AttitudeJaString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestAttitudePrintJa(t *testing.T) {
	expect := AttitudeJaString()
	actual := `[名]: [C, U]: [主な意味]: 態度、姿勢、考え方`

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}

// fair.go
func TestFairEnString(t *testing.T) {
	expect := "fair"
	actual := FairEnString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestFairJaString(t *testing.T) {
	expect := `[形]: [主な意味]: 1. 公平な、公正な 2. かなりの 3. まずまずの 4. (天候)晴れた、好天の
(また、[名]: [C]: [主な意味]: 博覧会、見本市、遊園地、説明会、慈善市など)`
	actual := FairJaString()

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
func TestFairPrintJa(t *testing.T) {
	expect := FairJaString()
	actual := `[形]: [主な意味]: 1. 公平な、公正な 2. かなりの 3. まずまずの 4. (天候)晴れた、好天の
(また、[名]: [C]: [主な意味]: 博覧会、見本市、遊園地、説明会、慈善市など)`

	if expect != actual {
		t.Errorf("%s != %s\n", expect, actual)
	}
}
