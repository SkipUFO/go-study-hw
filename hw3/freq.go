package main

import (
	"sort"
	"strings"
	"unicode"
)

type wordFreq struct {
	word  string
	count int
}

// TopWordsSlice function calculates top count frequency words in text.
// using slice in function
func TopWordsSlice(text string, count int) []string {

	var result []string
	var wordsFreq []wordFreq

	words := strings.FieldsFunc(text, func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	})

	// Создаем slice с wordFreq, потом его отсортируем
	for _, word := range words {
		word = strings.ToLower(word)

		var contains bool
		for i, val := range wordsFreq {
			if val.word == word {
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
	for _, val := range wordsFreq {
		if c >= count {
			break
		}

		result = append(result, val.word)
		c++

	}

	return result
}

// TopWordsMap function calculates top count frequency words in text.
// using map in function
func TopWordsMap(text string, count int) []string {

	var result []string
	wordsFreq := map[string]int{}

	words := strings.FieldsFunc(text, func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	})

	// Создаем map со словами и частотой
	for _, word := range words {
		word = strings.ToLower(word)

		if _, ok := wordsFreq[word]; ok {
			wordsFreq[word] = wordsFreq[word] + 1
		} else {
			wordsFreq[word] = 1
		}
	}

	// Создаем из map - slice для того, чтобы упорядочить
	slice := []wordFreq{}

	for key, value := range wordsFreq {
		slice = append(slice, wordFreq{word: key, count: value})
	}

	sort.Slice(slice, func(i, j int) bool {
		return (slice[i].count > slice[j].count) || ((slice[i].count == slice[j].count) && (slice[i].word < slice[j].word))
	})

	c := 0
	for _, val := range slice {
		if c >= count {
			break
		}

		result = append(result, val.word)
		c++

	}

	return result
}
