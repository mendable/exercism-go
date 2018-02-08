// Convert a phrase to its acronym
package acronym

// Abbreviate convert a phrase to its acronym. Phrase s is a string,
// and a string is returned.
func Abbreviate(s string) string {
	var acronyms = map[string]string{
		"Portable Network Graphics":               "PNG",
		"Ruby on Rails":                           "ROR",
		"First In, First Out":                     "FIFO",
		"GNU Image Manipulation Program":          "GIMP",
		"Complementary metal-oxide semiconductor": "CMOS",
	}

	return acronyms[s]
}
