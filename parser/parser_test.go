package parser_test

import (
	"github.com/friendsofgo/board-kata/parser"
	"testing"
)

func TestParse(t *testing.T) {
	tests := map[string]struct{ message, expectedMessage string }{
		"message without special characters": {"Es un test", "Es un test"},
		"message with arroba": {"Es un test de @fogo ", "Es un test de <a href='https://fogower.com/fogo'>@fogo</a>" },
		"message with hashtag": {"Es un test de #fogo ", "Es un test de <a href='https://fogower.com/hash/fogo'>#fogo</a>" },
		"message with url": {"Es un test de (link: http://bit.ly/oferta-go) ", "Es un test de <a href='http://bit.ly/oferta-go'>http://bit.ly/oferta-go</a>" },
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := parser.Parse(tc.message)

			if got != tc.expectedMessage {
				t.Fatalf(
					"Wrong parser expected: %s, got: %s",
					tc.expectedMessage,
					got,
				)
			}
		})
	}
}

func TestParseFail(t *testing.T) {
	tests := map[string]struct{ message, expectedErrorMessage string }{
		"message empty": {"", "empty message: empty message received"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := parser.Parse(tc.message)

			if err.Error() != tc.expectedErrorMessage {
				t.Fatalf(
					"Wrong parser expected: %s, got: %s",
					tc.expectedErrorMessage,
					err.Error(),
				)
			}
		})
	}
}