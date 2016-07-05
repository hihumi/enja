package enja

import (
	"fmt"
)

type WordsListStruct struct {
	wordsListEnString, wordsListJaString, wordsList string
}

func WordsListEnString() string {
	var WordsListEn WordsListStruct
	WordsListEn.wordsListEnString = "words-list"

	return WordsListEn.wordsListEnString
}

func WordsListJaString() string {
	var WordsListJa WordsListStruct
	WordsListJa.wordsListJaString = `words list:
[a]:
  attitude,

[b]:

[c]:
  creed,

[d]:

[e]:
  equality,

[f]:
  fair,

[g]:
  gender,

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
  national,
  negative,

[o]:
  origin,
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

	return WordsListJa.wordsListJaString
}

func WordsListPrintJa() {
	var WordsList WordsListStruct
	WordsList.wordsList = WordsListJaString()

	fmt.Println(WordsList.wordsList)

	return
}
