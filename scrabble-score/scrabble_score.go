package scrabble

import "unicode"
import "strings"

var GroupedScores = map[int]string{
	1:  "AEIOULNRST",
	2:  "DG",
	3:  "BCMP",
	4:  "FHVWY",
	5:  "K",
	8:  "JX",
	10: "QZ",
}

var ItemizedScores = map[string]int{
	"A": 1,
	"E": 1,
	"I": 1,
	"O": 1,
	"U": 1,
	"L": 1,
	"N": 1,
	"R": 1,
	"S": 1,
	"T": 1,
	"D": 2,
	"G": 2,
	"B": 3,
	"C": 3,
	"M": 3,
	"P": 3,
	"F": 4,
	"H": 4,
	"V": 4,
	"W": 4,
	"Y": 4,
	"K": 5,
	"J": 8,
	"X": 8,
	"Q": 10,
	"Z": 10,
}

var useItemizedTables = false

func Score(word string) (accumulator int) {
	if useItemizedTables {
		return ScoreFromItemization(word)
	} else {
		return ScoreFromGrouped(word)
	}
}

func ScoreFromGrouped(word string) (accumulator int) {
	for _, r := range word {
		accumulator += ItemizedScores[strings.ToUpper(string(r))]
	}

	return
}

func ScoreFromItemization(word string) (accumulator int) {
	for _, r := range word {
		for score, runes := range GroupedScores {
			if strings.ContainsRune(runes, unicode.ToUpper(r)) {
				accumulator += score
			}
		}
	}

	return
}
