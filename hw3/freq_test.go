package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseType struct {
	text  string
	count int
	words []string
}

var testCases = []testCaseType{
	{text: `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`,
		count: 2,
		words: []string{"in", "ut"},
	},
	{text: `Все смешалось в доме Облонских. Жена узнала, что муж был в связи с бывшею в их доме француженкою-гувернанткой, и объявила мужу, что не может жить с ним в одном доме. Положение это продолжалось уже третий день и мучительно чувствовалось и самими супругами, и всеми членами семьи, и домочадцами. Все члены семьи и домочадцы чувствовали, что нет смысла в их сожительстве и что на каждом постоялом дворе случайно сошедшиеся люди более связаны между собой, чем они, члены семьи и домочадцы Облонских. Жена не выходила из своих комнат, мужа третий день не было дома. Дети бегали по всему дому, как потерянные; англичанка поссорилась с экономкой и написала записку приятельнице, прося приискать ей новое место; повар ушел вчера со двора, во время самого обеда; черная кухарка и кучер просили расчета.`,
		count: 10,
		words: []string{"и", "в", "что", "доме", "не", "с", "семьи", "все", "день", "домочадцы"},
	},
	{text: `in out.`,
		count: 10,
		words: []string{"in", "out"},
	},
}

func sliceEqual(s1 []string, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func TestTopWordsSlice(t *testing.T) {

	for _, testCase := range testCases {
		result := TopWordsSlice(testCase.text, testCase.count)

		assert.Equal(t, true, sliceEqual(testCase.words, result))
	}
}

func TestTopWordsMap(t *testing.T) {

	for _, testCase := range testCases {
		result := TopWordsMap(testCase.text, testCase.count)

		assert.Equal(t, true, sliceEqual(testCase.words, result))
	}
}

func BenchmarkTopWordsSlice(b *testing.B) {
	for _, testCase := range testCases {
		TopWordsSlice(testCase.text, testCase.count)
	}
}

func BenchmarkTopWordsMap(b *testing.B) {
	for _, testCase := range testCases {
		TopWordsMap(testCase.text, testCase.count)
	}
}
