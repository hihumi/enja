package enjas

import (
	"fmt"
)

type WordsListStruct struct {
	wordsListEnString, wordsListReturnString, wordsList string
}

func WordsListEnString() string {
	var WordsListEn WordsListStruct
	WordsListEn.wordsListEnString = "words-list"

	return WordsListEn.wordsListEnString
}

func WordsListReturnString() string {
	var WordsListReturn WordsListStruct
	WordsListReturn.wordsListReturnString = `words list:
[a]:
  attitude,
  awful,

[b]:

[c]:
  constitution,
  creed,

[d]:

[e]:
  equality,

[f]:
  fair,

[g]:
  gaze,
  gender,
  guarantee,

[h]:

[i]:
  individual,

[j]:

[k]:

[l]:
  lean,
  liberty,

[m]:
  my,
  maintain,

[n]:
  national,
  negative,

[o]:
  origin,
  outlook,

[p]:
  perfume,
  pillar,
  positive,

[q]:

[r]:
  remind,
  respect,

[s]:
  scent,
  subtle,
  statue,

[t]:

[u]:

[v]:

[w]:
  will,

[x]:

[y]:

[z]:`

	return WordsListReturn.wordsListReturnString
}

func WordsListPrintString() {
	var WordsList WordsListStruct
	WordsList.wordsList = WordsListReturnString()

	fmt.Println(WordsList.wordsList)

	return
}
