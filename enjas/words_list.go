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
  gender,
  guarantee,

[h]:

[i]:
  individual,

[j]:

[k]:

[l]:
  lean,

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
  pillar,
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

	return WordsListReturn.wordsListReturnString
}

func WordsListPrintString() {
	var WordsList WordsListStruct
	WordsList.wordsList = WordsListReturnString()

	fmt.Println(WordsList.wordsList)

	return
}
