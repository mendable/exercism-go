// Bob is a lackadaisical teenager. In conversation, his responses are very limited.
package bob

import "strings"

// Talk to Bob through through the Hey function, passing a remark as a string.
// Returns a string response.
//
// Bob answers 'Sure.' if you ask him a question. (question mark)
// He answers 'Whoa, chill out!' if you yell at him. (uppercase)
// He answers 'Calm down, I know what I'm doing!' if you yell a question at him. (question mark + uppercase)
// He says 'Fine. Be that way!' if you address him without actually saying anything. (empty string, spaces, tabs etc)
// He answers 'Whatever.' to anything else.
func Hey(remark string) string {
	trimmedRemark := strings.TrimSpace(remark)

	isQuestion := strings.HasSuffix(trimmedRemark, "?")
	isYell := (strings.ToUpper(remark) == remark) && !(strings.ToLower(remark) == remark)

	if isYell && isQuestion {
		return "Calm down, I know what I'm doing!"
	} else if isYell {
		return "Whoa, chill out!"
	} else if isQuestion {
		return "Sure."
	} else if len(trimmedRemark) == 0 {
		return "Fine. Be that way!"
	} else {
		return "Whatever."
	}
}
