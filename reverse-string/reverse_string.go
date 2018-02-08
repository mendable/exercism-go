package reverse

// import "fmt"

// robot
// [0][1][2][3][4]
// [r][o][b][o][t]
//
// i=0, j=4
// t,r
//
// define iterators, first rune + last rune positions.
// loop so long as iterators do not cross (first counter less than second counter) (equal is middle rune on unequal word)
// Swap pair on each iteration
// - take rune from upward counter, downward counter
// - assign into opposite rune positions
// -- downward counter rune now in upward counter position
// -- upward counter rune now in downward counter position
// increment upward counter position
// decrement downward counter position
// repeat loop
func String(s string) string {
	// fmt.Printf("Reversing '%v'\n", s)
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		// fmt.Printf("Iteration %d %d (%v <=> %v)\n", i, j, runes[i], runes[j])
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
