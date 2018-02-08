package isogram

import "fmt"
import "unicode"

func contains(seen []rune, r rune) bool {
	for _, el := range seen {
		if r == el {
			return true
		}
	}
	return false
}

// take runes, add to seen list, check each rune against seen list, returning false if seen previously
// case insensitive. Non-letters e.g hyphens are not tested.
//
// Alternative implementation: use a map with rune as key and count as value.
func IsIsogram(word string) bool {
	var seen []rune

	fmt.Printf("Testing word: '%v' \n", word)

	for _, r := range word {
		if unicode.IsLetter(r) {
			if contains(seen, unicode.ToUpper(r)) {
				return false
			} else {
				seen = append(seen, unicode.ToUpper(r))
			}
		}
	}

	return true
}
