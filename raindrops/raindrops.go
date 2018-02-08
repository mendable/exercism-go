package raindrops

import "fmt"
import "bytes"

// If the number has 3 as a factor, output 'Pling'.
// If the number has 5 as a factor, output 'Plang'.
// If the number has 7 as a factor, output 'Plong'.
// If the number does not have 3, 5, or 7 as a factor, just pass the number's digits straight through.
func Convert(num int) string {
	var output bytes.Buffer

	if num%3 == 0 {
		output.WriteString("Pling")
	}

	if num%5 == 0 {
		output.WriteString("Plang")
	}

	if num%7 == 0 {
		output.WriteString("Plong")
	}

	if output.Len() > 0 {
		return fmt.Sprintf("%s", output.String())
	} else {
		return fmt.Sprintf("%d", num)
	}
}
