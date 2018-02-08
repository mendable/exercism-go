package twofer

import "fmt"

// ShareWith returns the sharing string with the name included.
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}

	return fmt.Sprintf("One for %v, one for me.", name)
}
