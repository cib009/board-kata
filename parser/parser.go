package parser

import (
	"regexp"
	"strings"
)

const (
	baseURL    = "https://fogower.com/"
	hashTagURL = "https://fogower.com/hash/"
)

// Parse fomating the text into a valid output text
func Parse(msg string) (output string) {

	words := strings.Fields(msg)

	for _, word :=range words {

		if word == "(link:" {
			continue
		}

		urlRegex := regexp.MustCompile(`(http|ftp|https):\/\/([\w\-]+(?:(?:\.[\w\-]+)+))([\w\-\.,@?^=%&amp;:/\+#]*[\w\-\@?^=%&amp;/\+#])?`)
		hashRegex := regexp.MustCompile(`(?:#)([a-zA-Z\d]+)?`)
		arrobaRegex := regexp.MustCompile(`(?:@)([a-zA-Z\d]+)?`)
		specialChars := regexp.MustCompile(`(?:\?)?(?:,)?(?:\))?`)

		if urlRegex.Match([]byte(word)) {
			word = specialChars.ReplaceAllString(word, "")
			output += " <a href='"+ word + "'>" + word + "</a>"
		} else if hashRegex.Match([]byte(word)) {
			word = specialChars.ReplaceAllString(word, "")
			output += " <a href='"+hashTagURL+strings.TrimPrefix( word,"#") + "'>" + word + "</a>"
		} else if arrobaRegex.Match([]byte(word)) {
			word = specialChars.ReplaceAllString(word, "")
			output += " <a href='"+baseURL+strings.TrimPrefix(word, "@") + "'>" + word + "</a>"
		} else {
			output += " " + word
		}
	}

	return output
}
