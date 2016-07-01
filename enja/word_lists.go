package enja

import (
	"fmt"
)

type WordListsStruct struct {
	wordListsEnString, wordListsJaString, wordLists string
}

func WordListsEnString() string {
	var WordListsEn WordListsStruct
	WordListsEn.wordListsEnString = "word-lists"

	return WordListsEn.wordListsEnString
}

func WordListsJaString() string {
	var WordListsJa WordListsStruct
	WordListsJa.wordListsJaString = `word lists:
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
  national,
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

	return WordListsJa.wordListsJaString
}

func WordListsPrintJa() {
	var WordLists WordListsStruct
	WordLists.wordLists = WordListsJaString()

	fmt.Println(WordLists.wordLists)

	return
}
