package main

import (
	"sort"
	"strings"
)

type wordFreq struct {
	word  string
	count int
}

// TopWords function calculates top count frequency words in text
func TopWords(text string, count int) []string {

	var result []string
	var wordsFreq []wordFreq

	words := strings.Split(text, " ")

	// Создаем slice с wordFreq, потом его отсортируем
	for _, word := range words {
		word = strings.ToLower(word)

		replacer := strings.NewReplacer(",", "", ".", "")
		word = replacer.Replace(word)

		var contains bool
		for i, val := range wordsFreq {
			if val.word == word {
				println("azaza", word)
				wordsFreq[i].count = wordsFreq[i].count + 1
				contains = true
				break
			}
		}

		if !contains {
			wordsFreq = append(wordsFreq, wordFreq{word: word, count: 1})
		}
	}

	sort.Slice(wordsFreq, func(i, j int) bool {
		return (wordsFreq[i].count > wordsFreq[j].count) || ((wordsFreq[i].count == wordsFreq[j].count) && (wordsFreq[i].word < wordsFreq[j].word))
	})

	c := 0
	for i := 0; i < len(wordsFreq); i++ {
		if c < count {
			result = append(result, wordsFreq[i].word)
			c = c + 1
		} else {
			break
		}
	}

	return result
}
